package feirarepository_test

import (
	"fmt"
	"testing"
	"utest/adapter/postgres/feirarepository"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupDelete() (pgxmock.PgxPoolIface) {	

	mock, _ := pgxmock.NewPool()

	return  mock
}
func TestDelete(t *testing.T) {


	mock := setupDelete()
	defer mock.Close()

	mock.ExpectQuery("delete from freira (.+)").
		WithArgs(1).RowsWillBeClosed()

	sut := feirarepository.New(mock)
	err := sut.Delete(1)

	require.NotNil(t, err)
}

func TestDelete_DBError(t *testing.T) {
	mock := setupDelete()
	defer mock.Close()

	mock.ExpectQuery("delete from feira where id (.+)").WithArgs(
		1       ,	
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := feirarepository.New(mock)
	err := sut.Delete(1)
	
	require.NotNil(t, err)
}

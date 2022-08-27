package feirarepository_test

import (
	"fmt"
	"testing"
	"utest/adapter/postgres/feirarepository"
	"utest/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupDelete() (dto.DeleteFeiraRequest, pgxmock.PgxPoolIface) {	
	fakeFeiraRequest := dto.DeleteFeiraRequest{}
	faker.FakeData(&fakeFeiraRequest)

	mock, _ := pgxmock.NewPool()

	return fakeFeiraRequest,  mock
}
func TestDelete(t *testing.T) {


	fakeFeiraRequest,  mock := setupDelete()
	defer mock.Close()

	mock.ExpectQuery("delete from freira (.+)").
		WithArgs(1).RowsWillBeClosed()

	sut := feirarepository.New(mock)
	err := sut.Delete(&fakeFeiraRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}


	require.NotNil(t, err)
}

func TestDelete_DBError(t *testing.T) {
	fakeFeiraRequest,mock := setupDelete()
	defer mock.Close()

	mock.ExpectQuery("delete from feira where id (.+)").WithArgs(
		fakeFeiraRequest.Id       ,	
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := feirarepository.New(mock)
	err := sut.Delete(&fakeFeiraRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.NotNil(t, err)
}

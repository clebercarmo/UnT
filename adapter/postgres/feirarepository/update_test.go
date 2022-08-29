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

func setupUpdate() (dto.UpdateFeiraRequest, pgxmock.PgxPoolIface) {	
	fakeFeiraRequest := dto.UpdateFeiraRequest{}
	faker.FakeData(&fakeFeiraRequest)

	mock, _ := pgxmock.NewPool()

	return fakeFeiraRequest, mock
}

func TestUpdate(t *testing.T) {
	fakeFeiraRequest, mock := setupUpdate()
	defer mock.Close()
	 mock.ExpectQuery("update feira set (.+)").WithArgs(
		1,
		fakeFeiraRequest.Long,
		fakeFeiraRequest.Lat,
		fakeFeiraRequest.SetCens,
		fakeFeiraRequest.AreaP,
		fakeFeiraRequest.CodDist,
		fakeFeiraRequest.Distrito,
		fakeFeiraRequest.CodSubPref,
		fakeFeiraRequest.SubPrere,
		fakeFeiraRequest.Regiao5,
		fakeFeiraRequest.Regiao8,
		fakeFeiraRequest.NomeFreira,
		fakeFeiraRequest.Registo,
		fakeFeiraRequest.Logradouro,
		fakeFeiraRequest.Numero,
		fakeFeiraRequest.Bairro,
		fakeFeiraRequest.Referencia,	
	).RowsWillBeClosed()


	
	
	sut := feirarepository.New(mock)
	feira, err := sut.Update(1, &fakeFeiraRequest)

	
	require.Nil(t, err)
	require.Equal(t, feira.Long, fakeFeiraRequest.Long)
	require.Equal(t, feira.Lat, fakeFeiraRequest.Lat)
	require.Equal(t, feira.SetCens, fakeFeiraRequest.SetCens)
	require.Equal(t, feira.AreaP, fakeFeiraRequest.AreaP)
	require.Equal(t, feira.CodDist, fakeFeiraRequest.CodDist)
	require.Equal(t, feira.Distrito, fakeFeiraRequest.Distrito)
	require.Equal(t, feira.CodSubPref, fakeFeiraRequest.CodSubPref)
	require.Equal(t, feira.SubPrere, fakeFeiraRequest.SubPrere)
	require.Equal(t, feira.Regiao5, fakeFeiraRequest.Regiao5)
	require.Equal(t, feira.Regiao8, fakeFeiraRequest.Regiao8)
	require.Equal(t, feira.NomeFreira, fakeFeiraRequest.NomeFreira)
	require.Equal(t, feira.Registo, fakeFeiraRequest.Registo)
	require.Equal(t, feira.Logradouro, fakeFeiraRequest.Logradouro)
	require.Equal(t, feira.Numero, fakeFeiraRequest.Numero)
	require.Equal(t, feira.Bairro, fakeFeiraRequest.Bairro)
	require.Equal(t, feira.Referencia, fakeFeiraRequest.Referencia)
}

func TestUpdate_DBError(t *testing.T) {
	fakeFeiraRequest,  mock := setupUpdate()
	defer mock.Close()

	mock.ExpectQuery("update freira set (.+)").WithArgs(
		fakeFeiraRequest.Long       ,  	
		fakeFeiraRequest.Lat        ,  	
		fakeFeiraRequest.SetCens    ,	
		fakeFeiraRequest.AreaP      ,	
		fakeFeiraRequest.CodDist    ,	
		fakeFeiraRequest.Distrito   ,	
		fakeFeiraRequest.CodSubPref ,	
		fakeFeiraRequest.SubPrere   ,	
		fakeFeiraRequest.Regiao5    ,	
		fakeFeiraRequest.Regiao8    ,	
		fakeFeiraRequest.NomeFreira ,	
		fakeFeiraRequest.Registo    ,	
		fakeFeiraRequest.Logradouro ,	
		fakeFeiraRequest.Numero     ,	
		fakeFeiraRequest.Bairro     ,	
		fakeFeiraRequest.Referencia ,	
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := feirarepository.New(mock)
	feira, err := sut.Update(1, &fakeFeiraRequest)


	require.NotNil(t, err)
	require.Nil(t, feira)
}

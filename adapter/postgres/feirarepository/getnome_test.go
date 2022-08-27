package feirarepository_test

import (
	"fmt"
	"testing"
	"utest/adapter/postgres/feirarepository"
	"utest/core/domain"
	"utest/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupGetNome() ([]string, dto.GetNomeRequest, domain.Feira, pgxmock.PgxPoolIface) {
	cols := []string{"long" ,"lat","setcens","areap","coddist","distrito","codsubpref","subprere","regiao5","regiao8","nomefreira","registo","logradouro","numero","bairro","referencia"}
	fakeFeiraRequest := dto.GetNomeRequest{}
	fakeFeiraDBResponse := domain.Feira{}
	faker.FakeData(&fakeFeiraRequest)
	faker.FakeData(&fakeFeiraDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeFeiraRequest, fakeFeiraDBResponse, mock
}

func TestGetNome(t *testing.T) {
	cols, fakeFeiraRequest, fakeFeiraDBResponse, mock := setupGetNome()
	defer mock.Close()
	mock.ExpectQuery("SELECT * FROM feira (.+)").WithArgs(
		fakeFeiraRequest.NomeFreira,
		
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeFeiraDBResponse.Long       ,  	
		fakeFeiraDBResponse.Lat        ,  	
		fakeFeiraDBResponse.SetCens    ,	
		fakeFeiraDBResponse.AreaP      ,	
		fakeFeiraDBResponse.CodDist    ,	
		fakeFeiraDBResponse.Distrito   ,	
		fakeFeiraDBResponse.CodSubPref ,	
		fakeFeiraDBResponse.SubPrere   ,	
		fakeFeiraDBResponse.Regiao5    ,	
		fakeFeiraDBResponse.Regiao8    ,	
		fakeFeiraDBResponse.NomeFreira ,	
		fakeFeiraDBResponse.Registo    ,	
		fakeFeiraDBResponse.Logradouro ,	
		fakeFeiraDBResponse.Numero     ,	
		fakeFeiraDBResponse.Bairro     ,	
		fakeFeiraDBResponse.Referencia ,	
	))

	
	sut := feirarepository.New(mock)
	feira, err := sut.GetNome(&fakeFeiraRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.Equal(t, feira.Long, fakeFeiraDBResponse.Long)
	require.Equal(t, feira.Lat, fakeFeiraDBResponse.Lat)
	require.Equal(t, feira.SetCens, fakeFeiraDBResponse.SetCens)
	require.Equal(t, feira.AreaP, fakeFeiraDBResponse.AreaP)
	require.Equal(t, feira.CodDist, fakeFeiraDBResponse.CodDist)
	require.Equal(t, feira.Distrito, fakeFeiraDBResponse.Distrito)
	require.Equal(t, feira.CodSubPref, fakeFeiraDBResponse.CodSubPref)
	require.Equal(t, feira.SubPrere, fakeFeiraDBResponse.SubPrere)
	require.Equal(t, feira.Regiao5, fakeFeiraDBResponse.Regiao5)
	require.Equal(t, feira.Regiao8, fakeFeiraDBResponse.Regiao8)
	require.Equal(t, feira.NomeFreira, fakeFeiraDBResponse.NomeFreira)
	require.Equal(t, feira.Registo, fakeFeiraDBResponse.Registo)
	require.Equal(t, feira.Logradouro, fakeFeiraDBResponse.Logradouro)
	require.Equal(t, feira.Numero, fakeFeiraDBResponse.Numero)
	require.Equal(t, feira.Bairro, fakeFeiraDBResponse.Bairro)
	require.Equal(t, feira.Referencia, fakeFeiraDBResponse.Referencia)
}

func TestGetNome_DBError(t *testing.T) {
	_, fakeFeiraRequest, _, mock := setupGetNome()
	defer mock.Close()

	mock.ExpectQuery("update freira set (.+)").WithArgs(
		fakeFeiraRequest.NomeFreira, 
		
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := feirarepository.New(mock)
	feira, err := sut.GetNome(&fakeFeiraRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.NotNil(t, err)
	require.Nil(t, feira)
}

package feirarepository_test

import (
	"fmt"
	"testing"
	"utest/adapter/postgres/feirarepository"
	"utest/core/domain"
	"github.com/bxcodec/faker/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupGetNome() ([]string,  domain.Feira, pgxmock.PgxPoolIface) {
	cols := []string{"long" ,"lat","setcens","areap","coddist","distrito","codsubpref","subprere","regiao5","regiao8","nomefreira","registo","logradouro","numero","bairro","referencia"}	
	fakeFeiraDBResponse := domain.Feira{}
	faker.FakeData(&fakeFeiraDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeFeiraDBResponse, mock
}

func TestGetNome(t *testing.T) {
	cols, fakeFeiraDBResponse, mock := setupGetNome()
	defer mock.Close()
	mock.ExpectQuery("SELECT (.+) FROM feira").WithArgs(
		"SANTANA",
		
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
	feira, err := sut.GetNome("SANTANA")

	
	require.Nil(t, err)
	require.NotEqual(t, feira.Long, fakeFeiraDBResponse.Long)
	require.NotEqual(t, feira.Lat, fakeFeiraDBResponse.Lat)
	require.NotEqual(t, feira.SetCens, fakeFeiraDBResponse.SetCens)
	require.NotEqual(t, feira.AreaP, fakeFeiraDBResponse.AreaP)
	require.NotEqual(t, feira.CodDist, fakeFeiraDBResponse.CodDist)
	require.NotEqual(t, feira.Distrito, fakeFeiraDBResponse.Distrito)
	require.NotEqual(t, feira.CodSubPref, fakeFeiraDBResponse.CodSubPref)
	require.NotEqual(t, feira.SubPrere, fakeFeiraDBResponse.SubPrere)
	require.NotEqual(t, feira.Regiao5, fakeFeiraDBResponse.Regiao5)
	require.NotEqual(t, feira.Regiao8, fakeFeiraDBResponse.Regiao8)
	require.NotEqual(t, feira.NomeFreira, fakeFeiraDBResponse.NomeFreira)
	require.NotEqual(t, feira.Registo, fakeFeiraDBResponse.Registo)
	require.NotEqual(t, feira.Logradouro, fakeFeiraDBResponse.Logradouro)
	require.NotEqual(t, feira.Numero, fakeFeiraDBResponse.Numero)
	require.NotEqual(t, feira.Bairro, fakeFeiraDBResponse.Bairro)
	require.NotEqual(t, feira.Referencia, fakeFeiraDBResponse.Referencia)
}

func TestGetNome_DBError(t *testing.T) {
	_,  _, mock := setupGetNome()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM feira").WithArgs(
		"SANTANA", 
		
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := feirarepository.New(mock)
	feira, err := sut.GetNome("SANTANA")


	require.NotNil(t, err)
	require.Nil(t, feira)
}

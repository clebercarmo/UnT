package dto_test

import (
	"encoding/json"
	"strings"
	"testing"
	"utest/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)


func TestFromJSONCreateFeiraRequest(t *testing.T) {
	fakeItem := dto.CreateFeiraRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONCreateFeiraRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.Id, fakeItem.Id)
	require.Equal(t, itemRequest.Lat, fakeItem.Lat)
	require.Equal(t, itemRequest.Long, fakeItem.Long)
	require.Equal(t, itemRequest.SetCens, fakeItem.SetCens)
	require.Equal(t, itemRequest.AreaP, fakeItem.AreaP)
	require.Equal(t, itemRequest.CodDist, fakeItem.CodDist)
	require.Equal(t, itemRequest.Distrito, fakeItem.Distrito)
	require.Equal(t, itemRequest.CodSubPref, fakeItem.CodSubPref)
	require.Equal(t, itemRequest.SubPrere, fakeItem.SubPrere)
	require.Equal(t, itemRequest.Regiao5, fakeItem.Regiao5)
	require.Equal(t, itemRequest.Regiao8, fakeItem.Regiao8)
	require.Equal(t, itemRequest.NomeFreira, fakeItem.NomeFreira)
	require.Equal(t, itemRequest.Registo, fakeItem.Registo)
	require.Equal(t, itemRequest.Logradouro, fakeItem.Logradouro)
	require.Equal(t, itemRequest.Numero, fakeItem.Numero)
	require.Equal(t, itemRequest.Bairro, fakeItem.Bairro)
	require.Equal(t, itemRequest.Referencia, fakeItem.Referencia)
}

func TestFromJSONCreateFeiraRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateFeiraRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}


func TestFromJSONUpdateFeiraRequest(t *testing.T) {

	fakeItem := dto.UpdateFeiraRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONUpdateFeiraRequest(strings.NewReader(string(json)))

	require.Nil(t, err)	
	require.Equal(t, itemRequest.Lat, fakeItem.Lat)
	require.Equal(t, itemRequest.Long, fakeItem.Long)
	require.Equal(t, itemRequest.SetCens, fakeItem.SetCens)
	require.Equal(t, itemRequest.AreaP, fakeItem.AreaP)
	require.Equal(t, itemRequest.CodDist, fakeItem.CodDist)
	require.Equal(t, itemRequest.Distrito, fakeItem.Distrito)
	require.Equal(t, itemRequest.CodSubPref, fakeItem.CodSubPref)
	require.Equal(t, itemRequest.SubPrere, fakeItem.SubPrere)
	require.Equal(t, itemRequest.Regiao5, fakeItem.Regiao5)
	require.Equal(t, itemRequest.Regiao8, fakeItem.Regiao8)
	require.Equal(t, itemRequest.NomeFreira, fakeItem.NomeFreira)
	require.Equal(t, itemRequest.Registo, fakeItem.Registo)
	require.Equal(t, itemRequest.Logradouro, fakeItem.Logradouro)
	require.Equal(t, itemRequest.Numero, fakeItem.Numero)
	require.Equal(t, itemRequest.Bairro, fakeItem.Bairro)
	require.Equal(t, itemRequest.Referencia, fakeItem.Referencia)

}


func TestFromJSONUpdateFeiraRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONUpdateFeiraRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}


func TestFromJSONDeleteFeiraRequest(t *testing.T) {

	fakeItem := dto.DeleteFeiraRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONDeleteFeiraRequest(strings.NewReader(string(json)))
	require.Nil(t, err)	
	require.Equal(t, itemRequest.Id, fakeItem.Id)

}

func TestFromJSONDeleteFeiraRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONDeleteFeiraRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}

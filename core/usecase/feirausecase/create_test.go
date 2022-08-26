package feirausecase_test

import (
	"fmt"
	"testing"
	"utest/core/domain"
	"utest/core/domain/mocks"
	"utest/core/dto"
	"utest/core/usecase/feirausecase"

	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	fakeRequestFeira := dto.CreateFeiraRequest{}
	fakeDBFeira := domain.Feira{}
	faker.FakeData(&fakeRequestFeira)
	faker.FakeData(&fakeDBFeira)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFeiraRepository := mocks.NewMockFeiraRepository(mockCtrl)
	mockFeiraRepository.EXPECT().Create(&fakeRequestFeira).Return(&fakeDBFeira, nil)

	sut := feirausecase.New(mockFeiraRepository)
	feira, err := sut.Create(&fakeRequestFeira)

	require.Nil(t, err)
	require.NotEmpty(t, feira.ID)
	require.Equal(t, feira.Long, fakeDBFeira.Long)
	require.Equal(t, feira.Lat, fakeDBFeira.Lat)
	require.Equal(t, feira.SetCens, fakeDBFeira.SetCens)
	require.Equal(t, feira.AreaP, fakeDBFeira.AreaP)
	require.Equal(t, feira.CodDist, fakeDBFeira.CodDist)
	require.Equal(t, feira.Distrito, fakeDBFeira.Distrito)
	require.Equal(t, feira.CodSubPref, fakeDBFeira.CodSubPref)
	require.Equal(t, feira.SubPrere, fakeDBFeira.SubPrere)
	require.Equal(t, feira.Regiao5, fakeDBFeira.Regiao5)
	require.Equal(t, feira.Regiao8, fakeDBFeira.Regiao8)
	require.Equal(t, feira.NomeFreira, fakeDBFeira.NomeFreira)
	require.Equal(t, feira.Registo, fakeDBFeira.Registo)
	require.Equal(t, feira.Logradouro, fakeDBFeira.Logradouro)
	require.Equal(t, feira.Numero, fakeDBFeira.Numero)
	require.Equal(t, feira.Bairro, fakeDBFeira.Bairro)
	require.Equal(t, feira.Referencia, fakeDBFeira.Referencia)
}

func TestCreate_Error(t *testing.T) {
	fakeRequestFeira := dto.CreateFeiraRequest{}
	faker.FakeData(&fakeRequestFeira)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFeiraRepository := mocks.NewMockFeiraRepository(mockCtrl)
	mockFeiraRepository.EXPECT().Create(&fakeRequestFeira).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := feirausecase.New(mockFeiraRepository)
	feira, err := sut.Create(&fakeRequestFeira)

	require.NotNil(t, err)
	require.Nil(t, feira)
}

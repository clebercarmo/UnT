package feirausecase_test

import (
	"fmt"
	"testing"
	"utest/core/dto"
	"utest/core/domain"
	"utest/core/domain/mocks"
	"utest/core/usecase/feirausecase"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	fakeRequestFeira := dto.DeleteFeiraRequest{}
	fakeDBFeira := domain.Feira{}
	faker.FakeData(&fakeRequestFeira)
	faker.FakeData(&fakeDBFeira)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFeiraRepository := mocks.NewMockFeiraRepository(mockCtrl)
	mockFeiraRepository.EXPECT().Delete(&fakeRequestFeira).Return(nil)

	sut := feirausecase.New(mockFeiraRepository)
	err := sut.Delete(&fakeRequestFeira)

	require.Nil(t, err)
	
}

func TestDelete_Error(t *testing.T) {
	fakeRequestFeira := dto.DeleteFeiraRequest{}
	faker.FakeData(&fakeRequestFeira)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFeiraRepository := mocks.NewMockFeiraRepository(mockCtrl)
	mockFeiraRepository.EXPECT().Delete(&fakeRequestFeira).Return(fmt.Errorf("ANY ERROR"))

	sut := feirausecase.New(mockFeiraRepository)
	err := sut.Delete(&fakeRequestFeira)

	require.NotNil(t, err)
}
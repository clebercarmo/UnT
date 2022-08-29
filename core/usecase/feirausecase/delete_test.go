package feirausecase_test

import (
	"fmt"
	"testing"
	"utest/core/domain"
	"utest/core/domain/mocks"
	"utest/core/usecase/feirausecase"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	fakeDBFeira := domain.Feira{}
	faker.FakeData(&fakeDBFeira)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFeiraRepository := mocks.NewMockFeiraRepository(mockCtrl)
	mockFeiraRepository.EXPECT().Delete(0).Return(nil)

	sut := feirausecase.New(mockFeiraRepository)
	err := sut.Delete(0)

	require.Nil(t, err)
	
}

func TestDelete_Error(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFeiraRepository := mocks.NewMockFeiraRepository(mockCtrl)
	mockFeiraRepository.EXPECT().Delete(0).Return(fmt.Errorf("ANY ERROR"))

	sut := feirausecase.New(mockFeiraRepository)
	err := sut.Delete(0)

	require.NotNil(t, err)
}
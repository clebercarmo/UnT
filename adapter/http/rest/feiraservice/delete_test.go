package feiraservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"utest/adapter/http/rest/feiraservice"
	"utest/core/domain/mocks"
	"github.com/golang/mock/gomock"
)

func setupDelete(t *testing.T) (*gomock.Controller) {	

	mockCtrl := gomock.NewController(t)

	return  mockCtrl
}

func TestDelete(t *testing.T) {
	mock := setupDelete(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().Delete(1).Return(nil)

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/feira/1", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Delete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestDelete_PorductError(t *testing.T) {
	mock := setupDelete(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().Delete(1).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/feira/1", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Delete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
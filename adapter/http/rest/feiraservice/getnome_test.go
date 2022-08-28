package feiraservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"utest/adapter/http/rest/feiraservice"
	"utest/core/domain"
	"utest/core/domain/mocks"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
)

func setupGetNome(t *testing.T) (domain.Feira, *gomock.Controller) {
	fakeFeira := domain.Feira{}
	faker.FakeData(&fakeFeira)

	mockCtrl := gomock.NewController(t)

	return  fakeFeira, mockCtrl
}

func TestGetNome(t *testing.T) {
	fakeFeira, mock := setupGetNome(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().GetNome("feira distrital").Return(&fakeFeira, nil)

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/feira", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("nome", "feira distrital")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.GetNome(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestGetNome_GetNomeError(t *testing.T) {
	 _, mock := setupGetNome(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().GetNome("feira distrital").Return(nil, fmt.Errorf("ANY ERROR"))

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/feira", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("nome", "feira distrital")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.GetNome(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
package feiraservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"utest/adapter/http/rest/feiraservice"
	"utest/core/domain"
	"utest/core/domain/mocks"
	"utest/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
)

func setupCreate(t *testing.T) (dto.CreateFeiraRequest, domain.Feira, *gomock.Controller) {
	fakeFeiraRequest := dto.CreateFeiraRequest{}
	fakeFeira := domain.Feira{}
	faker.FakeData(&fakeFeiraRequest)
	faker.FakeData(&fakeFeira)

	mockCtrl := gomock.NewController(t)

	return fakeFeiraRequest, fakeFeira, mockCtrl
}

func TestCreate(t *testing.T) {
	fakeFeiraRequest, fakeFeira, mock := setupCreate(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().Create(&fakeFeiraRequest).Return(&fakeFeira, nil)

	sut := feiraservice.New(mockFeiraUseCase)

	payload, _ := json.Marshal(fakeFeiraRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/feira", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	_, _, mock := setupCreate(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/feira", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_PorductError(t *testing.T) {
	fakeFeiraRequest, _, mock := setupCreate(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().Create(&fakeFeiraRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := feiraservice.New(mockFeiraUseCase)

	payload, _ := json.Marshal(fakeFeiraRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/feira", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
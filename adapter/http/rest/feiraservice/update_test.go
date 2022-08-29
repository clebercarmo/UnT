package feiraservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"utest/adapter/http/rest/feiraservice"
	"utest/core/domain"
	"utest/core/domain/mocks"
	"utest/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
)

func setupUpdate(t *testing.T) (dto.UpdateFeiraRequest, domain.Feira, *gomock.Controller) {
	fakePaginationRequestParams := dto.UpdateFeiraRequest{
		Long       : "121",
		Lat        : "435435",
		SetCens    : "555",
		AreaP      : "fdsgfsd",
		CodDist    :"vvvvv",
		Distrito   : "bbjhnh",
		CodSubPref : "jkuuuu",
		SubPrere   : "aaasfa",
		Regiao5    : "jjjjj",
		Regiao8    : "hhhhhh",
		NomeFreira : "jkkk",
		Registo    : "nnnn",
		Logradouro : "nmnvmvm",
		Numero     : "zzzzzz",
		Bairro     : "cccc",
		Referencia : "iiiii",
	}
	fakeFeira := domain.Feira{}
	faker.FakeData(&fakeFeira)

	mockCtrl := gomock.NewController(t)

	return fakePaginationRequestParams, fakeFeira, mockCtrl
}

func TestUpdate(t *testing.T) {
	fakePaginationRequestParams, fakeFeira, mock := setupUpdate(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().Update(&fakePaginationRequestParams).Return(&fakeFeira, nil)

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/feira/1", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()


}

func TestUpdate_FeiraError(t *testing.T) {
	fakePaginationRequestParams, _, mock := setupUpdate(t)
	defer mock.Finish()
	mockFeiraUseCase := mocks.NewMockFeiraUseCase(mock)
	mockFeiraUseCase.EXPECT().Update(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := feiraservice.New(mockFeiraUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/feira/1", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
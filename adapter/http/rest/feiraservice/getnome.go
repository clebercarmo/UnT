package feiraservice

import (
	"encoding/json"
	"net/http"
	"utest/core/dto"
)

// @Summary Obtem um registro de feira a partir de um nome
// @Description Obtem uma feira usando o nome da feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Feira
// @Router /api/feiras/{nome_feira}/nomes [get]
func (service service) GetNome(response http.ResponseWriter, request *http.Request) {
	feiraRequest, err := dto.FromJSONGetNameRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	feira, err := service.usecase.GetNome(feiraRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(feira)

}
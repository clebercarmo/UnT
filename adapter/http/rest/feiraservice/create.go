package feiraservice

import (
	"encoding/json"
	"net/http"
	"utest/core/dto"
)

// @Summary Cria um registro de feira
// @Description Cria uma nova feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Param feira body dto.CreateFeiraRequest true "feira"
// @Success 200 {object} domain.Feira
// @Router /feira [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {
	feiraRequest, err := dto.FromJSONCreateFeiraRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	feira, err := service.usecase.Create(feiraRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(feira)
}
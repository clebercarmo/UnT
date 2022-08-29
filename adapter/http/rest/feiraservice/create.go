package feiraservice

import (
	"encoding/json"
	"net/http"
	"utest/core/dto"
	"utest/shared"
)

// @Summary Cria um registro de feira
// @Description Cria uma nova feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Param feira body dto.CreateFeiraRequest true "feira"
// @Success 200 {object} domain.Feira
// @Router /api/feiras [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {
	feiraRequest, err := dto.FromJSONCreateFeiraRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	shared.SetLog("Info", "Criando feira " + feiraRequest.NomeFreira)
	feira, err := service.usecase.Create(feiraRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		shared.SetLog("Error", err.Error())
		return
	}
	
	shared.SetLog("Info", "Feira " + feiraRequest.NomeFreira + "Criada com Sucesso!")
	json.NewEncoder(response).Encode(feira)
}
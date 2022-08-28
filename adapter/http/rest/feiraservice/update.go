package feiraservice

import (
	"encoding/json"
	"net/http"
	"utest/core/dto"
)

// @Summary Altera um registro de feira
// @Description Altera uma nova feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Param feira body dto.UpdateFeiraRequest true "feira"
// @Success 200 {object} domain.Feira
// @Router /feira/1 [put]
func (service service) Update(response http.ResponseWriter, request *http.Request) {
	feiraRequest, err := dto.FromJSONUpdateFeiraRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	feira, err := service.usecase.Update(feiraRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(feira)
}
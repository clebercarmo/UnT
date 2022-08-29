package feiraservice

import (
	"encoding/json"
	"net/http"
	"strconv"
	"utest/core/dto"
	"utest/shared"
	"github.com/gorilla/mux"
)

// @Summary Altera um registro de feira
// @Description Altera uma nova feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Param feira body dto.UpdateFeiraRequest true "feira"
// @Success 200 {object} domain.Feira
// @Router /api/feiras/{id} [put]
func (service service) Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	feiraRequest, err := dto.FromJSONUpdateFeiraRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		shared.SetLog("Error", err.Error())
		return
	}

	feira, err := service.usecase.Update(id, feiraRequest)

	if err != nil {
		response.WriteHeader(404)
		var body = Body{Detail: err.Error()}
		json.NewEncoder(response).Encode(body)
		shared.SetLog("Error", err.Error())
		return
	}

	json.NewEncoder(response).Encode(feira)
}
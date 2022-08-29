package feiraservice

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


// @Summary Apaga um registro de feira
// @Description Apaga uma feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Feira
// @Router /api/feiras/{id} [delete]
func (service service) Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])


	err := service.usecase.Delete(id)

	if err != nil {
		response.WriteHeader(404)
		var body = Body{Detail: err.Error()}
		json.NewEncoder(response).Encode(body)
		return
	}

	var body = Body{Detail: "Excluido com sucesso"}

	json.NewEncoder(response).Encode(body)

}
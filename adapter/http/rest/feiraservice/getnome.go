package feiraservice

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)



// @Summary Obtem um registro de feira a partir de um nome
// @Description Obtem uma feira usando o nome da feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Feira
// @Router /api/feiras/{nome_feira}/nomes [get]
func (service service) GetNome(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	

	feira, err := service.usecase.GetNome(vars["nome_feira"])


	if feira.ID == 0 {
		response.WriteHeader(404)
		var body = Body{Detail: "Registro não existe"}
		json.NewEncoder(response).Encode(body)
		return
	}

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}


	json.NewEncoder(response).Encode(feira)

}
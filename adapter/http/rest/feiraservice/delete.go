package feiraservice

import (
	"net/http"
	"utest/core/dto"
)

// @Summary Apaga um registro de feira
// @Description Apaga uma feira
// @Tags feira
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Feira
// @Router /feira/1 [delete]
func (service service) Delete(response http.ResponseWriter, request *http.Request) {
	feiraRequest, err := dto.FromJSONDeleteFeiraRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	err = service.usecase.Delete(feiraRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

}
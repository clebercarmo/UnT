package feirarepository

import (
	"utest/core/domain"
	"utest/core/dto"
)

func (repository repository) Update(feiraRequest *dto.UpdateFeiraRequest)  (*domain.Feira, error){

	var feira domain.Feira
	repository.db.First(&feira, feiraRequest.Id) // find product with integer primary key

	feira.Long = feiraRequest.Long
	feira.Lat = feiraRequest.Lat
	feira.SetCens = feiraRequest.SetCens
	feira.AreaP = feiraRequest.AreaP
	feira.CodDist = feiraRequest.CodDist
	feira.Distrito = feiraRequest.Distrito
	feira.CodSubPref = feiraRequest.CodSubPref
	feira.SubPrere = feiraRequest.SubPrere
	feira.Regiao5 = feiraRequest.Regiao5
	feira.Regiao8 = feiraRequest.Regiao8
	feira.NomeFreira = feiraRequest.NomeFreira
	feira.Registo = feiraRequest.Registo
	feira.Logradouro = feiraRequest.Logradouro
	feira.Numero = feiraRequest.Numero
	feira.Bairro = feiraRequest.Bairro
	feira.Referencia = feiraRequest.Referencia

	repository.db.Save(&feira)

	return &feira, nil

}

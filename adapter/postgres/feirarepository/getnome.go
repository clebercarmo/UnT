package feirarepository

import (
	"utest/core/domain"
	"utest/core/dto"
)

func (repository repository) GetNome(feiraRequest *dto.GetNomeRequest)  (*domain.Feira, error){

	var feira domain.Feira
	repository.db.Find(&feira, domain.Feira{NomeFreira: feiraRequest.NomeFreira})

	return &feira, nil

}

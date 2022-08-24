package feirausecase

import (
	"utest/core/dto"
	"utest/core/domain"
)


func (usecase usecase) GetNome(feiraCreateRequest *dto.GetNomeRequest) (*domain.Feira, error) {
	feira, err := usecase.repository.GetNome(feiraCreateRequest)

	if err != nil {
		return nil, err
	}

	return feira, nil
}
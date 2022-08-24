package feirausecase

import (
	"utest/core/dto"
	"utest/core/domain"
)


func (usecase usecase) Create(feiraCreateRequest *dto.CreateFeiraRequest) (*domain.Feira, error) {
	feira, err := usecase.repository.Create(feiraCreateRequest)

	if err != nil {
		return nil, err
	}

	return feira, nil
}
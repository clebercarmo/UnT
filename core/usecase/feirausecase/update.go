package feirausecase

import (
	"utest/core/dto"
	"utest/core/domain"
)

func (usecase usecase) Update(id int,  feiraUpdateRequest *dto.UpdateFeiraRequest) (*domain.Feira, error) {
	feira, err := usecase.repository.Update(id, feiraUpdateRequest)

	if err != nil {
		return nil, err
	}

	return feira, nil
}
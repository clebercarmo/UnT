package feirausecase

import (
	"utest/core/domain"
)


func (usecase usecase) GetNome(nome string) (*domain.Feira, error) {
	feira, err := usecase.repository.GetNome(nome)

	if err != nil {
		return nil, err
	}

	return feira, nil
}
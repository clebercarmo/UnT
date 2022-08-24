package feirausecase

import "utest/core/domain"

type usecase struct {
	repository domain.FeiraRepository
}

func New(repository domain.FeiraRepository) domain.FeiraUseCase {
	return &usecase{
		repository: repository,
	}
}

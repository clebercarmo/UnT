package feirausecase

import "utest/core/domain"

type usercase struct {
	repository domain.FeiraRepository
}

func Init(repository domain.FeiraRepository) domain.FeiraUseCase {
	return &usercase{
		repository: repository,
	}
}

package feiraservice

import "utest/core/domain"

type Body struct{

	Detail string `json:"detail"`
}

type service struct {
	usecase domain.FeiraUseCase
}

// Retorna o contrato para implementar feiraservice
func New(usecase domain.FeiraUseCase) domain.FeiraService {
	return &service{
		usecase: usecase,
	}
}
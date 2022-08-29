package di

import (	
	"utest/adapter/http/rest/feiraservice"
	"utest/adapter/postgres"
	"utest/adapter/postgres/feirarepository"
	"utest/core/domain"
	"utest/core/usecase/feirausecase"
)

// ConfigFeiraDI retorna uma FeiraService abstracao com as dependencias injetadas
func ConfigFeiraDI(conn postgres.PoolInterface) domain.FeiraService {
	feiraRepository := feirarepository.New(conn)
	feiraUseCase := feirausecase.New(feiraRepository)
	feiraService := feiraservice.New(feiraUseCase)

	return feiraService
}
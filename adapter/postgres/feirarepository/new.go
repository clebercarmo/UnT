package feirarepository

import (
	"utest/adapter/postgres"
	"utest/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.FeiraRepository {
	return &repository{
		db: db,
	}
}
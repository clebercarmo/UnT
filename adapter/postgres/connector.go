package postgres

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"utest/core/domain"
)


type PoolInterface interface {
	Close()
	Create(model interface{})  *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Save(model interface{})  *gorm.DB 
	First(dest interface{}, conds ...interface{}) *gorm.DB 
	Delete(value interface{}, conds ...interface{}) *gorm.DB 
	
}

// GetConnection retorna uma conexao do banco sqllite
func GetConnection() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	
	if err != nil {
		return nil, err
	}
	return db, nil

}

// RunMigrations roda a criação e carga dos dados
func RunMigrations() error{
	db, err := GetConnection();
	if err != nil {
		return  err
	} 
	db.AutoMigrate(&domain.Feira{})
	return nil
}
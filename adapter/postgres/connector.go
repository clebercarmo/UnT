package postgres

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetConnection retorna uma conexao do banco sqllite
func GetConnection() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}

// RunMigrations roda a criação e carga dos dados
func RunMigrations() {
	db, err := GetConnection();
	if err != nil {
		return  err
	} 
	db.AutoMigrate(&Feira{})
	return nil
}
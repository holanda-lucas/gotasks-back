package database

import (
	"log"

	"github.com/holanda-lucas/gotasks-back/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectWithDatabase() {
	DB, err = gorm.Open(sqlite.Open("database.db"))

	if err != nil {
		log.Panic("Não foi possível se conectar com o banco de dados.")
	}

	// Criando as tabelas caso ainda não tenham sido criadas
	DB.Table("users").AutoMigrate(&models.User{})
	DB.Table("tasks").AutoMigrate(&models.Task{})
	DB.Table("blacklisted_tokens").AutoMigrate(&models.BlacklistedToken{})
}
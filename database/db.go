package database

import (
	"fmt"
	"github.com/huzairuje/chatat_backend_engineer/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func CreateDB(config *util.Config) *gorm.DB {
	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Errorf("failed to open database: %v", err)
	}
	SetConnection(db)
	return db
}

//GetConnection : Get Available Connection
func GetConnection() *gorm.DB {
	return conn
}

//SetConnection : Set Available Connection
func SetConnection(connection *gorm.DB) {
	conn = connection
}

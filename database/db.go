package database

import (
	"database/sql"
	"fmt"
	"github.com/huzairuje/chatat_backend_engineer/util"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var conn *sql.DB

func CreateDB(config *util.Config) *sql.DB {
	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName

	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Errorf("failed to open database: %v", err)
	}
	SetConnection(db)
	return db
}

//GetConnection : Get Available Connection
func GetConnection() *sql.DB {
	return conn
}

//SetConnection : Set Available Connection
func SetConnection(connection *sql.DB) {
	conn = connection
}

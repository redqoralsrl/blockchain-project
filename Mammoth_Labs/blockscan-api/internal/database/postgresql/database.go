package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"

	"blockscan-go/internal/config"
	"fmt"
	"log"
	"time"
)

func ConnectDatabase(config *config.Config) (db *sql.DB) {
	user := config.DBUser
	password := config.DBPassword
	dbname := config.DBName
	host := config.DBHost
	port := config.DBPort

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetMaxOpenConns(400)

	if err != nil {
		log.Panic(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
		return nil
	}
	return db
}

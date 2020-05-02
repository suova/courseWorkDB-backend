package queries

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mysql"
)

var (
	db *sqlx.DB
)

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sqlx.Connect("postgres",psqlInfo)
	//"postgres://docker:docker@127.0.0.1:5432/docker?sslmode=disable"

	if err != nil {
		fmt.Print(err)
		log.Panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Print(err)
		log.Panic(err)
	}

	doMigrate()
}

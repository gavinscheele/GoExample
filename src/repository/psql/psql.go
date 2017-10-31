package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PsqlDb struct {
	dbHost   string
	dbName   string
	username string
	password string
}

// Construct : Instantiates a new psql connection object
func NewDatabase(dbHost, dbName, username, password string) *PsqlDb {
	return &PsqlDb{
		dbHost:   dbHost,
		dbName:   dbName,
		username: username,
		password: password,
	}
}

func (db *PsqlDb) Open() (*sql.DB, error) {
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", db.dbHost, db.username, db.password, db.dbName)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
	}
	return conn, err
}

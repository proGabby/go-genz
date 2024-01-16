package postgressDbImpl

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type PostgresDBImpl struct {
	connStr string
}

func (postgresDBImpl *PostgresDBImpl) InitDatabase() (*sql.DB, error) {

	connString, ok := os.LookupEnv("DB_CONNECTION_STRING")

	postgresDBImpl.connStr = connString

	if !ok {
		log.Println("DB_CONNECTION_STRING variable not set")
	}
	if postgresDBImpl.connStr == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable not set")
	}

	// Connect to the database
	db, err := sql.Open("postgres", postgresDBImpl.connStr)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")

	return db, nil

}
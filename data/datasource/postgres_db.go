package postgressDatasource

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/proGabby/4genz/domain/entity"
)

type PostgresDBStore struct {
	DB *sql.DB
}

func NewPostgresDBStore(db *sql.DB) *PostgresDBStore {
	return &PostgresDBStore{
		DB: db,
	}
}

func InitDatabase() (*sql.DB, error) {

	connString, ok := os.LookupEnv("DB_CONNECTION_STRING")

	if !ok {
		log.Println("DB_CONNECTION_STRING variable not set")
	}
	if connString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable not set")
	}

	// Connect to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Printf("error opening db: %v\n", err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")

	return db, nil

}

func (db *PostgresDBStore) RegisterUser(name, email, profileImageUrl string, hashedPassword []byte) (*entity.User, error) {

	var userID int
	query := "INSERT INTO users(name, email, profile_image_url, password, is_verified, token_version,email_otp,forgot_password_otp,forget_password_otp_expiry_time ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err := db.DB.QueryRow(query, name, email, profileImageUrl, hashedPassword, false, 0, "", "", 0).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	createdUser := entity.NewUser(userID, name, email, profileImageUrl, false)

	return createdUser, nil
}

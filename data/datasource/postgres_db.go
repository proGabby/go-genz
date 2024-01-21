package postgressDatasource

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/proGabby/4genz/data/dto"
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

func (db *PostgresDBStore) UpdateUserImage(userId int, profileImageUrl string) (*dto.UserResponse, error) {

	var userResDto dto.UserResponse
	query := "UPDATE users SET profile_image_url = $2 WHERE id = $1 RETURNING id, name, email,profile_image_url"
	err := db.DB.QueryRow(query, userId, profileImageUrl).Scan(&userResDto.Id, &userResDto.Name, &userResDto.Email, &userResDto.ProfileImageUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &userResDto, nil
}

func (db *PostgresDBStore) VerifyUserCredentials(email string) (*entity.User, error) {

	var userRes entity.User
	query := "SELECT id, name, password, email, profile_image_url FROM users WHERE email = $1"
	err := db.DB.QueryRow(query, email).Scan(&userRes.Id, &userRes.Name, &userRes.Password, &userRes.Email, &userRes.ProfileImageUrl)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &userRes, nil
}

func (db *PostgresDBStore) GetUserByID(userId int) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, profile_image_url FROM users WHERE id = $1"
	err := db.DB.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Email, &user.ProfileImageUrl)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

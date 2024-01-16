package dto

import "database/sql"

type UserStore struct {
	DB *sql.DB
}
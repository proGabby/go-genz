package db_repository

import "database/sql"

type DatabaseRepository interface {
	InitDatabase() (*sql.DB, error);
}
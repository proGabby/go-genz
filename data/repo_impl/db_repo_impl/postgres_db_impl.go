package postgressDbImpl

import (
	"database/sql"

	postgressDatasource "github.com/proGabby/4genz/data/datasource"
)

type PostgresDBImpl struct {
	psql postgressDatasource.PostgresUserDBStore
}

func (p *PostgresDBImpl) InitDatabase() (*sql.DB, error) {
	return p.InitDatabase()
}

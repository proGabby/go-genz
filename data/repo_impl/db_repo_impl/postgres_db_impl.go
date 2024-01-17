package postgressDbImpl

import (
	"database/sql"

	postgressDatasource "github.com/proGabby/4genz/data/datasource"
)

type PostgresDBImpl struct {
	psql postgressDatasource.PostgresDBStore
}

func (p *PostgresDBImpl) InitDatabase() (*sql.DB, error) {
	return p.InitDatabase()
}



package data

import (
	"fmt"

	"github.com/Ressorrrrra/Test-Task/internal/pkg/config"
	"github.com/restream/reindexer"
)

type Database struct {
	Connection *reindexer.Reindexer
}

func CreateDb(cfg config.Config) (*Database, error) {
	var connectionString string
	db := &Database{}

	connectionString = fmt.Sprintf("cproto://%s:%s/%s",
		cfg.Db.Hostname,
		cfg.Db.Port,
		cfg.Db.Database)

	db.Connection = reindexer.NewReindex(connectionString, reindexer.WithCreateDBIfMissing())
	err := db.Connection.Ping()
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}

}

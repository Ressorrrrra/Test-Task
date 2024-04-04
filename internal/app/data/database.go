
import (
	"github.com/restream/reindexer"
	"errors"
	"fmt"
	"./internal/pkg/config/config"
)

type Database struct {
	Connection *reindexer.Reindexer
}

func CreateDb(config config.Config) (*Database, error) {
	var connectionString string
	db = &Database

	connectionString = fmt.Sprintf("cproto://%s:%s/%s",
config.Database.Hostname,
config.Database.Port,
config.Database.Database)

db.Connection = reindexer.NewReindexer(connectionString, reindexer.WithCreateDBIfMissing()) 
err := d.Connection.Ping()
if err != nil {
	return nil, err
}
else {
	return db, nil
}

}

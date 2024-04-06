package order

import (
	"github.com/Ressorrrrra/Test-Task/internal/app/data"
	"github.com/restream/reindexer"
)

type Order struct {
	ID        int
	Items     []*Item
	Cost      int
	OrderedAt int64
}

type Item struct {
	ID          int
	Name        string
	Description string
	Price       int
}

type Repository struct {
	Db *data.Database
}

func New(db *data.Database) (repos *Repository) {
	repos = &Repository{Db: db}
	return
}

func (r *Repository) Get() (orders []*Order, err error) {

	if namespaceErr := r.Db.Connection.OpenNamespace("Orders", &reindexer.NamespaceOptions{}, Order{}); namespaceErr != nil {
		return orders, namespaceErr
	}

	query := r.Db.Connection.Query("Orders").ReqTotal()
	qr := query.Exec()

	if execErr := qr.Error(); execErr != nil {
		return orders, execErr
	}

	for qr.Next() {
		item := qr.Object().(*Order)
		orders = append(orders, item)
	}

	return
}

func (r Repository) Create(order Order) error {

	if namespaceErr := r.Db.Connection.OpenNamespace("Orders", &reindexer.NamespaceOptions{}, Order{}); namespaceErr != nil {
		return namespaceErr
	}

	if upsertErr := r.Db.Connection.Upsert("Orders", (&order)); upsertErr != nil {
		return upsertErr
	}

	return nil
}

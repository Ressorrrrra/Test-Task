package order

import (
	"errors"
	"log"

	"github.com/Ressorrrrra/Test-Task/internal/app/data"
	"github.com/restream/reindexer"
)

type Order struct {
	ID        int `reindex:"id,,pk"`
	Items     []*Item
	Cost      int   `reindex:"cost"`
	OrderedAt int64 `reindex:"orderedAt"`
	Sort      int   `reindex:"sort"`
}

type Item struct {
	Name        string `reindex:"name"`
	Description string `reindex:"description"`
	Price       int    `reindex:"price"`
}

type Repository struct {
	Db *data.Database
}

func New(db *data.Database) (repos *Repository) {
	repos = &Repository{Db: db}

	return
}

func (r *Repository) Get() (orders []*Order, err error) {
	namespaceErr := r.Db.Connection.OpenNamespace("Orders", reindexer.DefaultNamespaceOptions(), Order{})
	defer r.Db.Connection.CloseNamespace("Orders")
	if namespaceErr != nil {
		return orders, namespaceErr
	}

	query := r.Db.Connection.Query("Orders").Sort("sort", true).ReqTotal()
	qr := query.Exec()
	defer qr.Close()
	if execErr := qr.Error(); execErr != nil {
		return orders, execErr
	}

	for qr.Next() {
		log.Println("appending")
		item := qr.Object().(*Order)
		orders = append(orders, item)
	}

	return
}

func (r *Repository) GetById(id int) (*Order, error) {
	namespaceErr := r.Db.Connection.OpenNamespace("Orders", reindexer.DefaultNamespaceOptions(), Order{})
	defer r.Db.Connection.CloseNamespace("Orders")
	if namespaceErr != nil {
		return nil, namespaceErr
	}

	order, found := r.Db.Connection.Query("Orders").Where("id", reindexer.EQ, id).Get()
	if found {
		return order.(*Order), nil
	} else {
		return nil, errors.New("object wasn't found")
	}
}

func (r *Repository) Create(order Order) error {

	namespaceErr := r.Db.Connection.OpenNamespace("Orders", reindexer.DefaultNamespaceOptions(), Order{})
	defer r.Db.Connection.CloseNamespace("Orders")
	if namespaceErr != nil {
		return namespaceErr
	}

	if _, upsertErr := r.Db.Connection.Insert("Orders", (&order), "id=serial()", "orderedAt=now()"); upsertErr != nil {
		return upsertErr
	}

	return nil
}

func (r *Repository) Update(order Order) error {

	namespaceErr := r.Db.Connection.OpenNamespace("Orders", reindexer.DefaultNamespaceOptions(), Order{})
	defer r.Db.Connection.CloseNamespace("Orders")
	if namespaceErr != nil {
		return namespaceErr
	}

	if _, upsertErr := r.Db.Connection.Update("Orders", (&order)); upsertErr != nil {
		return upsertErr
	}

	return nil
}

func (r *Repository) Delete(id int) error {

	namespaceErr := r.Db.Connection.OpenNamespace("Orders", reindexer.DefaultNamespaceOptions(), Order{})
	defer r.Db.Connection.CloseNamespace("Orders")
	if namespaceErr != nil {
		return namespaceErr
	}
	order, found := r.Db.Connection.Query("Orders").Where("id", reindexer.EQ, id).Get()
	if found {
		if upsertErr := r.Db.Connection.Delete("Orders", (order)); upsertErr != nil {
			return upsertErr
		}
	} else {
		return errors.New("object wasn't found")
	}
	return nil
}

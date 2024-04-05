package order

import "github.com/Ressorrrrra/Test-Task/internal/app/data"

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
	Db data.Database
}

func New() (repos *Repository) {

}

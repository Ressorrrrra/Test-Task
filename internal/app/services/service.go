package service

import "github.com/Ressorrrrra/Test-Task/internal/app/data/order"

type Service struct {
	repos *order.Repository
}

func New(rep *order.Repository) (s *Service) {
	s = &Service{repos: rep}
	return
}

func (s *Service) Get() (orders []*order.Order, err error) {
	orders, err = s.repos.Get()

	return
}

func (s *Service) Create(order order.Order, err error) {
	err = s.repos.Create(order)
}

func Update() {

}

func Delete() {

}

package service

import "github.com/Ressorrrrra/Test-Task/internal/app/data/order"

type Service struct {
	repos *order.Repository
}

func New(rep *order.Repository) (s *Service) {
	s = &Service{repos: rep}
	return
}

func (s *Service) Get() ([]*order.Order, error) {
	return s.repos.Get()
}

func (s *Service) GetById(id int) (*order.Order, error) {
	return s.repos.GetById(id)
}

func (s *Service) Create(order order.Order) error {
	return s.repos.Create(order)
}

func (s *Service) Update(order order.Order) error {
	return s.repos.Update(order)
}

func (s *Service) Delete(id int) error {
	return s.repos.Delete(id)
}

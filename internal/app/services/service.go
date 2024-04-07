package service

import (
	"log"
	"time"

	"github.com/Ressorrrrra/Test-Task/internal/app/data/order"
	"github.com/Ressorrrrra/Test-Task/internal/pkg/cache"
)

type DtoOrder struct {
	Items     []*order.Item
	Cost      int   `reindex:"cost"`
	OrderedAt int64 `reindex:"orderedAt"`
}

type Service struct {
	repos *order.Repository
	c     *cache.Cache
}

func New(rep *order.Repository) (s *Service) {
	c := cache.New(time.Minute, time.Minute*15)
	s = &Service{repos: rep, c: c}
	return
}

func ConvertToDto(orders []*order.Order) []*DtoOrder {
	resultChannel := make(chan []*DtoOrder)
	go func(_orders []*order.Order) {
		defer close(resultChannel)
		var dto []*DtoOrder
		for _, obj := range _orders {
			dtoObj := DtoOrder{
				Items:     obj.Items,
				Cost:      obj.Cost,
				OrderedAt: obj.OrderedAt,
			}
			dto = append(dto, &dtoObj)
		}
		resultChannel <- dto
	}(orders)

	dto := <-resultChannel
	return dto
}

func (s *Service) Get() ([]*DtoOrder, error) {

	orders, err := s.repos.Get()

	return ConvertToDto(orders), err
}

func (s *Service) GetPage(page, count int) ([]*DtoOrder, error) {
	orders, err := s.repos.GetOffset((page-1)*count, count)

	return ConvertToDto(orders), err
}

func (s *Service) GetById(id int) (*order.Order, error) {
	if item, found := s.c.Get(id); found {
		log.Println("Loaded from cache")
		return item.(*order.Order), nil
	}
	item, err := s.repos.GetById(id)
	if err == nil {
		s.c.Add(id, item, time.Minute*15)
		log.Println("Loaded in cache")
	}
	return item, err
}

func (s *Service) Create(order order.Order) error {
	return s.repos.Create(order)
}

func (s *Service) Update(order order.Order) error {
	if _, found := s.c.Get(order.ID); found {
		log.Println("Updated cache")
		s.c.Add(order.ID, order, time.Minute*15)
	}
	return s.repos.Update(order)
}

func (s *Service) Delete(id int) error {
	err := s.c.Delete(id)
	if err == nil {
		log.Println("Deleted from cache")
	}
	return s.repos.Delete(id)
}

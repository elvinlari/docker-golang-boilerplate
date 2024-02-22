package mock

import "github.com/elvinlari/docker-golang/internal/company/domain"

type Service struct {
	GetByIDFn      func(id int) (*domain.Company, error)
	GetByIDInvoked bool

	ListFn      func() ([]*domain.Company, error)
	ListInvoked bool

	CreateFn      func(t *domain.Company) (*domain.Company, error)
	CreateInvoked bool

	UpdateFn      func(t *domain.Company) (*domain.Company, error)
	UpdateInvoked bool

	DeleteFn      func(id int) error
	DeleteInvoked bool
}

func (s *Service) GetByID(id int) (*domain.Company, error) {
	s.GetByIDInvoked = true
	return s.GetByIDFn(id)
}

func (s *Service) List() ([]*domain.Company, error) {
	s.ListInvoked = true
	return s.ListFn()
}

func (s *Service) Create(t *domain.Company) (*domain.Company, error) {
	s.CreateInvoked = true
	return s.CreateFn(t)
}

func (s *Service) Update(t *domain.Company) (*domain.Company, error) {
	s.UpdateInvoked = true
	return s.UpdateFn(t)
}

func (s *Service) Delete(id int) error {
	s.DeleteInvoked = true
	return s.DeleteFn(id)
}

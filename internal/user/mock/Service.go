package mock

import "github.com/elvinlari/docker-golang/internal/user/domain"

type Service struct {
	GetByIDFn      func(id int) (*domain.User, error)
	GetByIDInvoked bool

	ListFn      func() ([]*domain.User, error)
	ListInvoked bool

	CreateFn      func(t *domain.User) (*domain.User, error)
	CreateInvoked bool

	UpdateFn      func(t *domain.User) (*domain.User, error)
	UpdateInvoked bool

	DeleteFn      func(id int) error
	DeleteInvoked bool
}

func (s *Service) GetByID(id int) (*domain.User, error) {
	s.GetByIDInvoked = true
	return s.GetByIDFn(id)
}

func (s *Service) List() ([]*domain.User, error) {
	s.ListInvoked = true
	return s.ListFn()
}

func (s *Service) Create(t *domain.User) (*domain.User, error) {
	s.CreateInvoked = true
	return s.CreateFn(t)
}

func (s *Service) Update(t *domain.User) (*domain.User, error) {
	s.UpdateInvoked = true
	return s.UpdateFn(t)
}

func (s *Service) Delete(id int) error {
	s.DeleteInvoked = true
	return s.DeleteFn(id)
}

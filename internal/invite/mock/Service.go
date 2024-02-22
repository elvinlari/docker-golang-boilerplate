package mock

import "github.com/elvinlari/docker-golang/internal/invite/domain"

type Service struct {
	GetByIDFn      func(id int) (*domain.Invite, error)
	GetByIDInvoked bool

	ListFn      func() ([]*domain.Invite, error)
	ListInvoked bool

	CreateFn      func(t *domain.Invite) (*domain.Invite, error)
	CreateInvoked bool

	UpdateFn      func(t *domain.Invite) (*domain.Invite, error)
	UpdateInvoked bool

	DeleteFn      func(id int) error
	DeleteInvoked bool
}

func (s *Service) GetByID(id int) (*domain.Invite, error) {
	s.GetByIDInvoked = true
	return s.GetByIDFn(id)
}

func (s *Service) List() ([]*domain.Invite, error) {
	s.ListInvoked = true
	return s.ListFn()
}

func (s *Service) Create(t *domain.Invite) (*domain.Invite, error) {
	s.CreateInvoked = true
	return s.CreateFn(t)
}

func (s *Service) Update(t *domain.Invite) (*domain.Invite, error) {
	s.UpdateInvoked = true
	return s.UpdateFn(t)
}

func (s *Service) Delete(id int) error {
	s.DeleteInvoked = true
	return s.DeleteFn(id)
}

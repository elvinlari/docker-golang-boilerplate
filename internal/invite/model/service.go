package model

import (
	"github.com/elvinlari/docker-golang/internal/invite/domain"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetByID(id int) (*domain.Invite, error) {
	var t domain.Invite
	tx := s.DB.First(&t, id)
	return &t, tx.Error
}

func (s *Service) List() ([]*domain.Invite, error) {
	var invites []*domain.Invite
	tx := s.DB.Find(&invites)
	return invites, tx.Error
}

func (s *Service) Create(t *domain.Invite) (*domain.Invite, error) {
	tx := s.DB.Create(t)
	return t, tx.Error
}

func (s *Service) Update(t *domain.Invite) (*domain.Invite, error) {
    var existingInvite domain.Invite
    if err := s.DB.First(&existingInvite, t.ID).Error; err != nil {
        return nil, err
    }
    if err := s.DB.Model(&existingInvite).Updates(t).Error; err != nil {
        return nil, err
    }
    return &existingInvite, nil
}

func (s *Service) Delete(id int) error {
	invite := &domain.Invite{ID: id}
	s.DB.Delete(&invite)
	return nil
}


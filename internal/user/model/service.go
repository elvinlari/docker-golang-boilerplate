package model

import (
	"github.com/elvinlari/docker-golang/internal/user/domain"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetByID(id int) (*domain.User, error) {
	var t domain.User
	tx := s.DB.First(&t, id)
	return &t, tx.Error
}

func (s *Service) List() ([]*domain.User, error) {
	var users []*domain.User
	tx := s.DB.Find(&users)
	return users, tx.Error
}

func (s *Service) Create(t *domain.User) (*domain.User, error) {
	tx := s.DB.Create(t)
	return t, tx.Error
}

func (s *Service) Update(t *domain.User) (*domain.User, error) {
    var existingUser domain.User
    if err := s.DB.First(&existingUser, t.ID).Error; err != nil {
        return nil, err
    }
    if err := s.DB.Model(&existingUser).Updates(t).Error; err != nil {
        return nil, err
    }
    return &existingUser, nil
}

func (s *Service) Delete(id int) error {
	user := &domain.User{ID: id}
	s.DB.Delete(&user)
	return nil
}


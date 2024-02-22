package model

import (
	"github.com/elvinlari/docker-golang/internal/company/domain"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetByID(id int) (*domain.Company, error) {
	var t domain.Company
	tx := s.DB.First(&t, id)
	return &t, tx.Error
}

func (s *Service) List() ([]*domain.Company, error) {
	var companies []*domain.Company
	tx := s.DB.Find(&companies)
	return companies, tx.Error
}

func (s *Service) Create(t *domain.Company) (*domain.Company, error) {
	tx := s.DB.Create(t)
	return t, tx.Error
}

func (s *Service) Update(t *domain.Company) (*domain.Company, error) {
    var existingCompany domain.Company
    if err := s.DB.First(&existingCompany, t.ID).Error; err != nil {
        return nil, err
    }
    if err := s.DB.Model(&existingCompany).Updates(t).Error; err != nil {
        return nil, err
    }
    return &existingCompany, nil
}

func (s *Service) Delete(id int) error {
	company := &domain.Company{ID: id}
	s.DB.Delete(&company)
	return nil
}


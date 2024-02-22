package domain

import (
	"time"
)

type Company struct {
	ID          int
	Name      	string
	Email       string
	Website 	string
	CountryCode int
	PhoneCode   int
	PhoneNumber string
	AddressLine1 string
	AddressLine2 string
	State    	string
	City    	string
	ZipCode     int
	Country     string
	Timezone 	string
	Status    	string
	Comment     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	GetByID(id int) (*Company, error)
	List() ([]*Company, error)
	Create(t *Company) (*Company, error)
	Update(t *Company) (*Company, error)
	Delete(id int) error
}
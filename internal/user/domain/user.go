package domain

import (
	"time"
)

type User struct {
	ID          int
	UUID      	string
	CompanyId   int
	Username 	string
	Email		string
	PhoneNumber string
	FirstName	string
	MiddleName	string
	LastName   	string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	GetByID(id int) (*User, error)
	List() ([]*User, error)
	Create(t *User) (*User, error)
	Update(t *User) (*User, error)
	Delete(id int) error
}
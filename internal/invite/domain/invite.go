package domain

import (
	"time"
)

type Invite struct {
	ID          int
	CompanyId   int
	Username    string
	Email       string
	FirstName   string
	LastName    string
	Enabled 	bool
	Status 		string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	GetByID(id int) (*Invite, error)
	List() ([]*Invite, error)
	Create(t *Invite) (*Invite, error)
	Update(t *Invite) (*Invite, error)
	Delete(id int) error
}
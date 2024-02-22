package http

import (

	"github.com/elvinlari/docker-golang/internal/user/domain"
)

// User is a struct with a subset of the fields of domain.User. It is used when
// User needs to be provided as an input for User creation. So it excludes
// auto-generated fields.
type User struct {
	UUID        string        `json:"uuid" example:"adefegewgfew"`
	CompanyId 	int        	  `json:"company_id" example:"34"`
	Username 	string        `json:"username" example:"username"`
	Email 		string        `json:"email" example:"email@website.com"`
	PhoneNumber	string        `json:"phone_number" example:"254712345678"`
	FirstName   string        `json:"first_name" example:"Micah"`
	MiddleName  string        `json:"middle_name" example:"Sam"`
	LastName	string        `json:"last_name" example:"Oduori"`
}

type Request struct {
	User *User `json:"user" binding:"required"`
}

func (t *User) httpToModel() *domain.User {
	return &domain.User{
		UUID:        t.UUID,
		CompanyId: 	 t.CompanyId,
		Username:    t.Username,
		Email: 		 t.Email,
		PhoneNumber: t.PhoneNumber,
		FirstName:   t.FirstName,
		MiddleName:  t.MiddleName,
		LastName:    t.LastName,
	}
}
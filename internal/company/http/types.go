package http

import (

	"github.com/elvinlari/docker-golang/internal/company/domain"
)

// Company is a struct with a subset of the fields of domain.Company. It is used when
// company needs to be provided as an input for company creation. So it excludes
// auto-generated fields.
type Company struct {
	Name        string        `json:"name" example:"my-company-1"`
	Email 		string        `json:"email" example:"my-company-1.email.com"`
	Website		string        `json:"website" example:"website.com"`
	CountryCode int           `json:"country_code" example:"254"`
	PhoneCode	int           `json:"phone_code" example:"254"`
	PhoneNumber	string        `json:"phone_number" example:"254712345678"`
	AddressLine1 string        `json:"address_line1" example:"2566 Dow ST."`
	AddressLine2 string        `json:"address_line2" example:"P.O. box 100"`
	State		string        `json:"state" example:"Nairobi"`
	City		string        `json:"city" example:"Nairobi"`
	ZipCode		int        	  `json:"zip_code" example:"00100" format:"int64"`
	Country		string        `json:"country" example:"Kenya"`
	Timezone	string        `json:"timezone" example:"UTC"`
	Status 		string        `json:"status" example:"Active"`
	Comment     string        `json:"comment" example:"Company comment"`
}

type Request struct {
	Company *Company `json:"company" binding:"required"`
}

func (t *Company) httpToModel() *domain.Company {
	return &domain.Company{
		Name:        t.Name,
		Email: 		 t.Email,
		Website:     t.Website,
		CountryCode: t.CountryCode,
		PhoneCode:   t.PhoneCode,
		PhoneNumber: t.PhoneNumber,
		AddressLine1: t.AddressLine1,
		AddressLine2: t.AddressLine2,
		State:       t.State,
		City:        t.City,
		ZipCode:     t.ZipCode,
		Country: 	 t.Country,
		Timezone:    t.Timezone,
		Status:      t.Status,
		Comment:     t.Comment,
	}
}
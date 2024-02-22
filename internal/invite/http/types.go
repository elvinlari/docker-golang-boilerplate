package http

import (

	"github.com/elvinlari/docker-golang/internal/invite/domain"
)

// Invite is a struct with a subset of the fields of domain.Invite. It is used when
// invite needs to be provided as an input for invite creation. So it excludes
// auto-generated fields.
type Invite struct {
	CompanyId 	int        	  `json:"company_id" example:"34"`
	Username 	string        `json:"username" example:"username"`
	Email 		string        `json:"email" example:"email@website.com"`
	FirstName   string        `json:"first_name" example:"Micah"`
	LastName	string        `json:"last_name" example:"Oduori"`
	Enabled     bool          `json:"enabled" example:"true"`
	Status      string        `json:"status" example:"pending_creation"`
}

type Request struct {
	Invite *Invite `json:"invite" binding:"required"`
}

func (t *Invite) httpToModel() *domain.Invite {
	return &domain.Invite{
		CompanyId: 	 t.CompanyId,
		Username:    t.Username,
		Email: 		 t.Email,
		FirstName:   t.FirstName,
		LastName:    t.LastName,
		Enabled: 	 t.Enabled,
		Status:  	 t.Status,
	}
}
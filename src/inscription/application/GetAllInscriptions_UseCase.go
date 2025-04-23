package application

import (
	"estsoftware/src/inscription/domain"
	"estsoftware/src/inscription/domain/entities"
)

type GetAllInscriptions struct {
	db domain.IInscription
}

func NewGetAllInscriptions(db domain.IInscription) *GetAllInscriptions {
	return &GetAllInscriptions{db: db}
}

func (uc *GetAllInscriptions) Execute() ([]entities.Inscription, error) {
	return uc.db.GetAll()
}

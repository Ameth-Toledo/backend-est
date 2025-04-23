package application

import (
	"estsoftware/src/inscription/domain"
	"estsoftware/src/inscription/domain/entities"
)

type CreateInscription struct {
	db domain.IInscription
}

func NewCreateInscription(db domain.IInscription) *CreateInscription {
	return &CreateInscription{db: db}
}

func (uc *CreateInscription) Execute(inscription entities.Inscription) (*entities.Inscription, error) {
	return uc.db.Save(inscription)
}

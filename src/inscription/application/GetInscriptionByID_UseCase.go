package application

import (
	"estsoftware/src/inscription/domain"
	"estsoftware/src/inscription/domain/entities"
	"fmt"
)

type GetInscriptionByID struct {
	db domain.IInscription
}

func NewGetInscriptionByID(db domain.IInscription) *GetInscriptionByID {
	return &GetInscriptionByID{db: db}
}

func (uc *GetInscriptionByID) Execute(id int) (*entities.Inscription, error) {
	inscription, err := uc.db.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("inscripción no encontrada: %v", err)
	}
	return inscription, nil
}

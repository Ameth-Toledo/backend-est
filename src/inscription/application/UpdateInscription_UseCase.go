package application

import (
	"estsoftware/src/inscription/domain"
	"estsoftware/src/inscription/domain/entities"
	"fmt"
)

type UpdateInscription struct {
	db domain.IInscription
}

func NewUpdateInscription(db domain.IInscription) *UpdateInscription {
	return &UpdateInscription{db: db}
}

func (uc *UpdateInscription) Execute(id int, inscription entities.Inscription) (*entities.Inscription, error) {
	existingInscription, err := uc.db.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("inscripción no encontrada: %v", err)
	}

	existingInscription.AlumnoID = inscription.AlumnoID
	existingInscription.CursoID = inscription.CursoID
	existingInscription.FechaInscripcion = inscription.FechaInscripcion

	updatedInscription, err := uc.db.Save(*existingInscription)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar inscripción: %v", err)
	}

	return updatedInscription, nil
}

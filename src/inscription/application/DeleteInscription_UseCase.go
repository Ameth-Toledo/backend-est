package application

import (
	"estsoftware/src/inscription/domain"
	"fmt"
)

type DeleteInscription struct {
	db domain.IInscription
}

func NewDeleteInscription(db domain.IInscription) *DeleteInscription {
	return &DeleteInscription{db: db}
}

func (uc *DeleteInscription) Execute(id int) error {
	err := uc.db.Delete(id)
	if err != nil {
		return fmt.Errorf("error al eliminar inscripción: %v", err)
	}
	return nil
}

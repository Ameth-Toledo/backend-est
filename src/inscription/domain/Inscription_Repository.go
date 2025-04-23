package domain

import "estsoftware/src/inscription/domain/entities"

type IInscription interface {
	Save(inscripcion entities.Inscription) (*entities.Inscription, error)
	GetAll() ([]entities.Inscription, error)
	GetByID(id int) (*entities.Inscription, error)
	Update(id int, inscripcion entities.Inscription) (*entities.Inscription, error)
	Delete(id int) error
}

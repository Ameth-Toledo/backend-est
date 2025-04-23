package domain

import "estsoftware/src/materials/domain/entities"

type IMaterial interface {
	Save(material entities.Material) (*entities.Material, error)
	GetAll() ([]entities.Material, error)
	GetById(id int) (*entities.Material, error)
	Update(material entities.Material) (*entities.Material, error)
	Delete(id int) error
}

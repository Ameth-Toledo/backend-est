package application

import (
	"estsoftware/src/materials/domain"
	"estsoftware/src/materials/domain/entities"
)

type CreateMaterial struct {
	db domain.IMaterial
}

func NewCreateMaterial(db domain.IMaterial) *CreateMaterial {
	return &CreateMaterial{db: db}
}

func (uc *CreateMaterial) Execute(material entities.Material) (*entities.Material, error) {
	return uc.db.Save(material)
}

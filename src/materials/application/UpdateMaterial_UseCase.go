package application

import (
	"estsoftware/src/materials/domain"
	"estsoftware/src/materials/domain/entities"
)

type UpdateMaterial struct {
	db domain.IMaterial
}

func NewUpdateMaterial(db domain.IMaterial) *UpdateMaterial {
	return &UpdateMaterial{db: db}
}

func (uc *UpdateMaterial) Execute(material entities.Material) (*entities.Material, error) {
	return uc.db.Update(material)
}

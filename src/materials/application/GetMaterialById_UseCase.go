package application

import (
	"estsoftware/src/materials/domain"
	"estsoftware/src/materials/domain/entities"
)

type GetMaterialById struct {
	db domain.IMaterial
}

func NewGetMaterialById(db domain.IMaterial) *GetMaterialById {
	return &GetMaterialById{db: db}
}

func (uc *GetMaterialById) Execute(id int) (*entities.Material, error) {
	return uc.db.GetById(id)
}

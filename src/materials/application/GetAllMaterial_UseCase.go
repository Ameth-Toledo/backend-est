package application

import (
	"estsoftware/src/materials/domain"
	"estsoftware/src/materials/domain/entities"
)

type GetAllMaterials struct {
	db domain.IMaterial
}

func NewGetAllMaterials(db domain.IMaterial) *GetAllMaterials {
	return &GetAllMaterials{db: db}
}

func (uc *GetAllMaterials) Execute() ([]entities.Material, error) {
	return uc.db.GetAll()
}

package application

import "estsoftware/src/materials/domain"

type DeleteMaterial struct {
	db domain.IMaterial
}

func NewDeleteMaterial(db domain.IMaterial) *DeleteMaterial {
	return &DeleteMaterial{db: db}
}

func (uc *DeleteMaterial) Execute(id int) error {
	return uc.db.Delete(id)
}

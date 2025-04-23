package application

import "estsoftware/src/modules/domain"

type DeleteModule struct {
	db domain.IModule
}

func NewDeleteModule(db domain.IModule) *DeleteModule {
	return &DeleteModule{db: db}
}

func (uc *DeleteModule) Execute(id int) error {
	return uc.db.Delete(id)
}

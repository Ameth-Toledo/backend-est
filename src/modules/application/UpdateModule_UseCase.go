package application

import (
	"estsoftware/src/modules/domain"
	"estsoftware/src/modules/domain/entities"
)

type UpdateModule struct {
	db domain.IModule
}

func NewUpdateModule(db domain.IModule) *UpdateModule {
	return &UpdateModule{db: db}
}

func (uc *UpdateModule) Execute(modulo entities.Module) (*entities.Module, error) {
	return uc.db.Update(modulo)
}

package application

import (
	"estsoftware/src/modules/domain"
	"estsoftware/src/modules/domain/entities"
)

type GetModuleById struct {
	db domain.IModule
}

func NewGetModuleById(db domain.IModule) *GetModuleById {
	return &GetModuleById{db: db}
}

func (uc *GetModuleById) Execute(id int) (*entities.Module, error) {
	return uc.db.GetById(id)
}

package application

import (
	"estsoftware/src/modules/domain"
	"estsoftware/src/modules/domain/entities"
)

type GetAllModules struct {
	db domain.IModule
}

func NewGetAllModules(db domain.IModule) *GetAllModules {
	return &GetAllModules{db: db}
}

func (uc *GetAllModules) Execute() ([]entities.Module, error) {
	return uc.db.GetAll()
}

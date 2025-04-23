package domain

import "estsoftware/src/modules/domain/entities"

type IModule interface {
	Save(module entities.Module) (*entities.Module, error)
	GetAll() ([]entities.Module, error)
	GetById(id int) (*entities.Module, error)
	Update(modulo entities.Module) (*entities.Module, error)
	Delete(id int) error
}

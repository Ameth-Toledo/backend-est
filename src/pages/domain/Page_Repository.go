package domain

import "estsoftware/src/pages/domain/entities"

type IPage interface {
	Save(page entities.Page) (*entities.Page, error)
	GetAll() ([]entities.Page, error)
	GetById(id int) (*entities.Page, error)
	Update(page entities.Page) (*entities.Page, error)
	Delete(id int) error
}

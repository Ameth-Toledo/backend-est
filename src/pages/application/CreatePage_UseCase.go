package application

import (
	"estsoftware/src/pages/domain"
	"estsoftware/src/pages/domain/entities"
)

type CreatePage struct {
	db domain.IPage
}

func NewCreatePage(db domain.IPage) *CreatePage {
	return &CreatePage{db: db}
}

func (uc *CreatePage) Execute(page entities.Page) (*entities.Page, error) {
	return uc.db.Save(page)
}

package application

import (
	"estsoftware/src/pages/domain"
	"estsoftware/src/pages/domain/entities"
)

type UpdatePage struct {
	db domain.IPage
}

func NewUpdatePage(db domain.IPage) *UpdatePage {
	return &UpdatePage{db: db}
}

func (u *UpdatePage) Execute(page entities.Page) (*entities.Page, error) {
	return u.db.Update(page)
}

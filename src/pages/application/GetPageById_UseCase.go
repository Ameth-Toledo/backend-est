package application

import (
	"estsoftware/src/pages/domain"
	"estsoftware/src/pages/domain/entities"
)

type GetPageById struct {
	db domain.IPage
}

func NewGetPageById(db domain.IPage) *GetPageById {
	return &GetPageById{db: db}
}

func (g *GetPageById) Execute(id int) (*entities.Page, error) {
	return g.db.GetById(id)
}

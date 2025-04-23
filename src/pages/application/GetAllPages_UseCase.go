package application

import (
	"estsoftware/src/pages/domain"
	"estsoftware/src/pages/domain/entities"
)

type GetAllPages struct {
	db domain.IPage
}

func NewGetAllPages(db domain.IPage) *GetAllPages {
	return &GetAllPages{db: db}
}

func (g *GetAllPages) Execute() ([]entities.Page, error) {
	return g.db.GetAll()
}

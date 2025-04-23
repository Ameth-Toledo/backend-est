package application

import "estsoftware/src/pages/domain"

type DeletePage struct {
	db domain.IPage
}

func NewDeletePage(db domain.IPage) *DeletePage {
	return &DeletePage{db: db}
}

func (d *DeletePage) Execute(id int) error {
	return d.db.Delete(id)
}

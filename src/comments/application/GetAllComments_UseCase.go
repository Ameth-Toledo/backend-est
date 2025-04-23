package application

import (
	"estsoftware/src/comments/domain"
	"estsoftware/src/comments/domain/entities"
)

type GetAllComments struct {
	db domain.IComment
}

func NewGetAllComments(db domain.IComment) *GetAllComments {
	return &GetAllComments{db: db}
}

func (uc *GetAllComments) Execute() ([]entities.Comment, error) {
	return uc.db.GetAll()
}

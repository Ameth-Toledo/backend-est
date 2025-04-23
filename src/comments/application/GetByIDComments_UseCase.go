package application

import (
	"estsoftware/src/comments/domain"
	"estsoftware/src/comments/domain/entities"
)

type GetCommentById struct {
	db domain.IComment
}

func NewGetCommentById(db domain.IComment) *GetCommentById {
	return &GetCommentById{db: db}
}

func (uc *GetCommentById) Execute(id int) (*entities.Comment, error) {
	return uc.db.GetById(id)
}

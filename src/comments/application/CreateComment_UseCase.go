package application

import (
	"estsoftware/src/comments/domain"
	"estsoftware/src/comments/domain/entities"
)

type CreateComment struct {
	db domain.IComment
}

func NewCreateComment(db domain.IComment) *CreateComment {
	return &CreateComment{db: db}
}

func (uc *CreateComment) Execute(comment entities.Comment) (*entities.Comment, error) {
	return uc.db.Save(comment)
}

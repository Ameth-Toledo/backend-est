package application

import (
	"estsoftware/src/comments/domain"
	"estsoftware/src/comments/domain/entities"
)

type UpdateComment struct {
	db domain.IComment
}

func NewUpdateComment(db domain.IComment) *UpdateComment {
	return &UpdateComment{db: db}
}

func (uc *UpdateComment) Execute(comment entities.Comment) (*entities.Comment, error) {
	return uc.db.Update(comment)
}

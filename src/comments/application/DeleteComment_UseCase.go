package application

import "estsoftware/src/comments/domain"

type DeleteComment struct {
	db domain.IComment
}

func NewDeleteComment(db domain.IComment) *DeleteComment {
	return &DeleteComment{db: db}
}

func (uc *DeleteComment) Execute(id int) error {
	return uc.db.Delete(id)
}

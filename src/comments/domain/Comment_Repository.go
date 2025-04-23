package domain

import "estsoftware/src/comments/domain/entities"

type IComment interface {
	Save(comment entities.Comment) (*entities.Comment, error)
	GetAll() ([]entities.Comment, error)
	GetById(id int) (*entities.Comment, error)
	Update(comment entities.Comment) (*entities.Comment, error)
	Delete(id int) error
}

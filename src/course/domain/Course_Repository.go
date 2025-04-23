package domain

import "estsoftware/src/course/domain/entities"

type ICourse interface {
	Save(course entities.Course) (*entities.Course, error)
	GetAll() ([]entities.Course, error)
	GetById(id int) (*entities.Course, error)
	GetByName(name string) ([]entities.Course, error)
	Update(entities.Course) error
	Delete(id int) error
}

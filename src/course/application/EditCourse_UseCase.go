package application

import (
	"estsoftware/src/course/domain"
	"estsoftware/src/course/domain/entities"
)

type UpdateCourse struct {
	db domain.ICourse
}

func NewUpdateCourse(db domain.ICourse) *UpdateCourse {
	return &UpdateCourse{db: db}
}

func (uc *UpdateCourse) Execute(course entities.Course) error {
	return uc.db.Update(course)
}

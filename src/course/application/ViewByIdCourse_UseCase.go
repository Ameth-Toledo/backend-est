package application

import (
	"estsoftware/src/course/domain"
	"estsoftware/src/course/domain/entities"
)

type GetCourseById struct {
	db domain.ICourse
}

func NewGetCourseById(db domain.ICourse) *GetCourseById {
	return &GetCourseById{db: db}
}

func (gcb *GetCourseById) Execute(id int) (*entities.Course, error) {
	return gcb.db.GetById(id)
}

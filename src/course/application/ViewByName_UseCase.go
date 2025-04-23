package application

import (
	"estsoftware/src/course/domain"
	"estsoftware/src/course/domain/entities"
)

type GetCourseByName struct {
	db domain.ICourse
}

func NewGetCourseByName(db domain.ICourse) *GetCourseByName {
	return &GetCourseByName{db: db}
}

func (gcn *GetCourseByName) Execute(name string) ([]entities.Course, error) {
	return gcn.db.GetByName(name)
}

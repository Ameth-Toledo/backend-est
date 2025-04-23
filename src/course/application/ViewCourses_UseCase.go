package application

import (
	"estsoftware/src/course/domain"
	"estsoftware/src/course/domain/entities"
)

type GetAllCourses struct {
	db domain.ICourse
}

func NewGetAllCourses(db domain.ICourse) *GetAllCourses {
	return &GetAllCourses{db: db}
}

func (gac *GetAllCourses) Execute() ([]entities.Course, error) {
	return gac.db.GetAll()
}

package application

import "estsoftware/src/course/domain"

type DeleteCourse struct {
	db domain.ICourse
}

func NewDeleteCourse(db domain.ICourse) *DeleteCourse {
	return &DeleteCourse{db: db}
}

func (dc *DeleteCourse) Execute(id int) error {
	return dc.db.Delete(id)
}

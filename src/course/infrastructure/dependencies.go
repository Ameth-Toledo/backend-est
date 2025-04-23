package infrastructure

import (
	"estsoftware/src/core"
	"estsoftware/src/course/application"
	"estsoftware/src/course/infrastructure/adapters"
	"estsoftware/src/course/infrastructure/controllers"
)

type DependenciesCourse struct {
	CreateCourseController    *controllers.CreateCourseController
	GetCoursesController      *controllers.GetAllCoursesController
	GetCourseByIdController   *controllers.GetCourseByIdController
	GetCourseByNameController *controllers.GetCourseByNameController
	UpdateCourseController    *controllers.UpdateCourseController
	DeleteCourseController    *controllers.DeleteCourseController
}

func InitCourses() *DependenciesCourse {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &DependenciesCourse{
		CreateCourseController:    controllers.NewCreateCourseController(application.NewCreateCourse(ps)),
		GetCoursesController:      controllers.NewGetAllCoursesController(application.NewGetAllCourses(ps)),
		GetCourseByIdController:   controllers.NewGetCourseByIdController(application.NewGetCourseById(ps)),
		GetCourseByNameController: controllers.NewGetCourseByNameController(application.NewGetCourseByName(ps)),
		UpdateCourseController:    controllers.NewUpdateCourseController(application.NewUpdateCourse(ps)),
		DeleteCourseController:    controllers.NewDeleteCourseController(application.NewDeleteCourse(ps)),
	}
}

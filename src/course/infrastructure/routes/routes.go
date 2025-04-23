package routes

import (
	"estsoftware/src/course/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureCoursesRoutes(r *gin.Engine,
	createCourseController *controllers.CreateCourseController,
	getCourseController *controllers.GetAllCoursesController,
	getByIdCourseController *controllers.GetCourseByIdController,
	getByNameCourseController *controllers.GetCourseByNameController,
	updateCourseController *controllers.UpdateCourseController,
	deleteCourseController *controllers.DeleteCourseController,
) {
	r.Static("/image_portada", "./image_portada")

	r.POST("/courses", createCourseController.Execute)
	r.GET("/courses", getCourseController.Execute)
	r.GET("/courses/id/:id", getByIdCourseController.Execute)
	r.GET("/courses/name/:name", getByNameCourseController.Execute)
	r.PUT("/courses/:id", updateCourseController.Execute)
	r.DELETE("/courses/:id", deleteCourseController.Execute)
}

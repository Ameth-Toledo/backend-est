package controllers

import (
	"estsoftware/src/course/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllCoursesController struct {
	useCase *application.GetAllCourses
}

func NewGetAllCoursesController(useCase *application.GetAllCourses) *GetAllCoursesController {
	return &GetAllCoursesController{useCase: useCase}
}

func (cc *GetAllCoursesController) Execute(context *gin.Context) {
	courses, err := cc.useCase.Execute()
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudieron obtener los cursos"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"courses": courses})
}

package controllers

import (
	"estsoftware/src/course/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCourseByNameController struct {
	useCase *application.GetCourseByName
}

func NewGetCourseByNameController(useCase *application.GetCourseByName) *GetCourseByNameController {
	return &GetCourseByNameController{useCase: useCase}
}

func (cc *GetCourseByNameController) Execute(context *gin.Context) {
	name := context.Param("name")
	if name == "" {
		context.JSON(400, gin.H{"error": "El nombre es obligatorio"})
		return
	}

	courses, err := cc.useCase.Execute(name)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudieron obtener los cursos"})
		return
	}

	if len(courses) == 0 {
		context.JSON(404, gin.H{"error": "No se encontraron cursos"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"courses": courses})
}

package controllers

import (
	"estsoftware/src/course/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetCourseByIdController struct {
	useCase *application.GetCourseById
}

func NewGetCourseByIdController(useCase *application.GetCourseById) *GetCourseByIdController {
	return &GetCourseByIdController{useCase: useCase}
}

func (cc *GetCourseByIdController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	course, err := cc.useCase.Execute(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo obtener el curso"})
		return
	}

	if course == nil {
		context.JSON(404, gin.H{"error": "Curso no encontrado"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"course": course})
}

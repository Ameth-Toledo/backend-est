package controllers

import (
	"estsoftware/src/course/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteCourseController struct {
	useCase *application.DeleteCourse
}

func NewDeleteCourseController(useCase *application.DeleteCourse) *DeleteCourseController {
	return &DeleteCourseController{useCase: useCase}
}

func (cc *DeleteCourseController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	if err := cc.useCase.Execute(id); err != nil {
		context.JSON(500, gin.H{"error": "No se pudo eliminar el curso"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "Curso eliminado"})
}

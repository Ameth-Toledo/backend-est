package controllers

import (
	"estsoftware/src/course/application"
	"estsoftware/src/course/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateCourseController struct {
	useCase *application.UpdateCourse
}

func NewUpdateCourseController(useCase *application.UpdateCourse) *UpdateCourseController {
	return &UpdateCourseController{useCase: useCase}
}

func (cc *UpdateCourseController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var course entities.Course
	if err := context.ShouldBindJSON(&course); err != nil {
		context.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	course.ID = int32(id)

	if err := cc.useCase.Execute(course); err != nil {
		context.JSON(500, gin.H{"error": "No se pudo actualizar el curso"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "Curso actualizado"})
}

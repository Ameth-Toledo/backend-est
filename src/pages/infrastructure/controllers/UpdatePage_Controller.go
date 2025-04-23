package controllers

import (
	"estsoftware/src/pages/application"
	"estsoftware/src/pages/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePageController struct {
	useCase *application.UpdatePage
}

func NewUpdatePageController(useCase *application.UpdatePage) *UpdatePageController {
	return &UpdatePageController{useCase: useCase}
}

func (u *UpdatePageController) Execute(context *gin.Context) {
	var page entities.Page
	if err := context.ShouldBindJSON(&page); err != nil {
		context.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	updatedPage, err := u.useCase.Execute(page)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo actualizar la página"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"page": updatedPage})
}

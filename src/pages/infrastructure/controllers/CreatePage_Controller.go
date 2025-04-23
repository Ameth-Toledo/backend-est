package controllers

import (
	"estsoftware/src/pages/application"
	"estsoftware/src/pages/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePageController struct {
	useCase *application.CreatePage
}

func NewCreatePageController(useCase *application.CreatePage) *CreatePageController {
	return &CreatePageController{useCase: useCase}
}

func (cc *CreatePageController) Execute(context *gin.Context) {
	var page entities.Page
	if err := context.ShouldBindJSON(&page); err != nil {
		context.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	createdPage, err := cc.useCase.Execute(page)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo crear la página"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"page": createdPage})
}

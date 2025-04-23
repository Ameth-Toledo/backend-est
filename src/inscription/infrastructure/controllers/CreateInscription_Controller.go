package controllers

import (
	"estsoftware/src/inscription/application"
	"estsoftware/src/inscription/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateInscriptionController struct {
	useCase *application.CreateInscription
}

func NewCreateInscripcionController(useCase *application.CreateInscription) *CreateInscriptionController {
	return &CreateInscriptionController{useCase: useCase}
}

func (c *CreateInscriptionController) Execute(ctx *gin.Context) {
	var input entities.Inscription

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	result, err := c.useCase.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

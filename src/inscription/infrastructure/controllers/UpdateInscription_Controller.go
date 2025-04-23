package controllers

import (
	"estsoftware/src/inscription/application"
	"estsoftware/src/inscription/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateInscriptionController struct {
	useCase *application.UpdateInscription
}

func NewUpdateInscriptionController(useCase *application.UpdateInscription) *UpdateInscriptionController {
	return &UpdateInscriptionController{useCase: useCase}
}

func (c *UpdateInscriptionController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input entities.Inscription
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	updatedInscription, err := c.useCase.Execute(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedInscription)
}

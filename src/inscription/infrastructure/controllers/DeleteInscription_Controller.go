package controllers

import (
	"estsoftware/src/inscription/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteInscriptionController struct {
	useCase *application.DeleteInscription
}

func NewDeleteInscriptionController(useCase *application.DeleteInscription) *DeleteInscriptionController {
	return &DeleteInscriptionController{useCase: useCase}
}

func (c *DeleteInscriptionController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Inscripción eliminada"})
}

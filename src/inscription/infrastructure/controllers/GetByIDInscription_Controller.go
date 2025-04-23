package controllers

import (
	"estsoftware/src/inscription/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetInscriptionByIDController struct {
	useCase *application.GetInscriptionByID
}

func NewGetInscriptionByIDController(useCase *application.GetInscriptionByID) *GetInscriptionByIDController {
	return &GetInscriptionByIDController{useCase: useCase}
}

func (c *GetInscriptionByIDController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	inscription, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inscription)
}

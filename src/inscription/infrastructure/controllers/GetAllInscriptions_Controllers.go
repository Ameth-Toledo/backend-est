package controllers

import (
	"estsoftware/src/inscription/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllInscriptionsController struct {
	useCase *application.GetAllInscriptions
}

func NewGetAllInscriptionsController(useCase *application.GetAllInscriptions) *GetAllInscriptionsController {
	return &GetAllInscriptionsController{useCase: useCase}
}

func (c *GetAllInscriptionsController) Execute(ctx *gin.Context) {
	inscriptions, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inscriptions)
}

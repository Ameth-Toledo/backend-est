package controllers

import (
	"estsoftware/src/pages/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllPagesController struct {
	useCase *application.GetAllPages
}

func NewGetAllPagesController(useCase *application.GetAllPages) *GetAllPagesController {
	return &GetAllPagesController{useCase: useCase}
}

func (g *GetAllPagesController) Execute(context *gin.Context) {
	pages, err := g.useCase.Execute()
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudieron obtener las páginas"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"pages": pages})
}

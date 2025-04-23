package controllers

import (
	"estsoftware/src/pages/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetPageByIdController struct {
	useCase *application.GetPageById
}

func NewGetPageByIdController(useCase *application.GetPageById) *GetPageByIdController {
	return &GetPageByIdController{useCase: useCase}
}

func (g *GetPageByIdController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID no válido"})
		return
	}

	page, err := g.useCase.Execute(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo obtener la página"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"page": page})
}

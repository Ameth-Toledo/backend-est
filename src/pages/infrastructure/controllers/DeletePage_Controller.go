package controllers

import (
	"estsoftware/src/pages/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeletePageController struct {
	useCase *application.DeletePage
}

func NewDeletePageController(useCase *application.DeletePage) *DeletePageController {
	return &DeletePageController{useCase: useCase}
}

func (d *DeletePageController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID no válido"})
		return
	}

	err = d.useCase.Execute(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo eliminar la página"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Página eliminada con éxito"})
}

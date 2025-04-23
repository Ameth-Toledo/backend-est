package controllers

import (
	"estsoftware/src/modules/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteModuleController struct {
	useCase *application.DeleteModule
}

func NewDeleteModuleController(useCase *application.DeleteModule) *DeleteModuleController {
	return &DeleteModuleController{useCase: useCase}
}

func (cc *DeleteModuleController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	err = cc.useCase.Execute(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo eliminar el módulo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Módulo eliminado con éxito"})
}

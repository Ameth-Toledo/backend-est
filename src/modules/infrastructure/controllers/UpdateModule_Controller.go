package controllers

import (
	"estsoftware/src/modules/application"
	"estsoftware/src/modules/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateModuleController struct {
	useCase *application.UpdateModule
}

func NewUpdateModuleController(useCase *application.UpdateModule) *UpdateModuleController {
	return &UpdateModuleController{useCase: useCase}
}

func (cc *UpdateModuleController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var modulo entities.Module
	if err := context.ShouldBindJSON(&modulo); err != nil {
		context.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	modulo.ID = int32(id)

	updatedModule, err := cc.useCase.Execute(modulo)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo actualizar el módulo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"module": updatedModule})
}

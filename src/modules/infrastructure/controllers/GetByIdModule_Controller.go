package controllers

import (
	"estsoftware/src/modules/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetModuleByIdController struct {
	useCase *application.GetModuleById
}

func NewGetModuleByIdController(useCase *application.GetModuleById) *GetModuleByIdController {
	return &GetModuleByIdController{useCase: useCase}
}

func (cc *GetModuleByIdController) Execute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	module, err := cc.useCase.Execute(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo obtener el módulo"})
		return
	}

	if module == nil {
		context.JSON(404, gin.H{"error": "Módulo no encontrado"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"module": module})
}

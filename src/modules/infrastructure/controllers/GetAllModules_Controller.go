package controllers

import (
	"estsoftware/src/modules/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllModulesController struct {
	useCase *application.GetAllModules
}

func NewGetAllModulesController(useCase *application.GetAllModules) *GetAllModulesController {
	return &GetAllModulesController{useCase: useCase}
}

func (cc *GetAllModulesController) Execute(context *gin.Context) {
	modules, err := cc.useCase.Execute()
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudieron obtener los módulos"})
		return
	}

	if len(modules) == 0 {
		context.JSON(404, gin.H{"error": "No se encontraron módulos"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"modules": modules})
}

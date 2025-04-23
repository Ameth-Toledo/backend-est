package controllers

import (
	"estsoftware/src/materials/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllMaterialsController struct {
	UseCase *application.GetAllMaterials
}

func NewGetAllMaterialsController(useCase *application.GetAllMaterials) *GetAllMaterialsController {
	return &GetAllMaterialsController{UseCase: useCase}
}

func (c *GetAllMaterialsController) Execute(ctx *gin.Context) {
	materials, err := c.UseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los materiales"})
		return
	}

	ctx.JSON(http.StatusOK, materials)
}

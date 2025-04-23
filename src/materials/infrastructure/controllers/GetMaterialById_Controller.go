package controllers

import (
	"estsoftware/src/materials/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetMaterialByIdController struct {
	UseCase *application.GetMaterialById
}

func NewGetMaterialByIdController(useCase *application.GetMaterialById) *GetMaterialByIdController {
	return &GetMaterialByIdController{UseCase: useCase}
}

func (c *GetMaterialByIdController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	material, err := c.UseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Material no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, material)
}

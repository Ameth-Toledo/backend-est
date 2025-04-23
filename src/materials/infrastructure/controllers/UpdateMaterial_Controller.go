package controllers

import (
	"estsoftware/src/materials/application"
	"estsoftware/src/materials/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateMaterialController struct {
	UseCase *application.UpdateMaterial
}

func NewUpdateMaterialController(useCase *application.UpdateMaterial) *UpdateMaterialController {
	return &UpdateMaterialController{UseCase: useCase}
}

func (c *UpdateMaterialController) Execute(ctx *gin.Context) {
	var material entities.Material
	if err := ctx.ShouldBindJSON(&material); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	updated, err := c.UseCase.Execute(material)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el material"})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

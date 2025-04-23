package controllers

import (
	"estsoftware/src/materials/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteMaterialController struct {
	UseCase *application.DeleteMaterial
}

func NewDeleteMaterialController(useCase *application.DeleteMaterial) *DeleteMaterialController {
	return &DeleteMaterialController{UseCase: useCase}
}

func (c *DeleteMaterialController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.UseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el material"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Material eliminado con éxito"})
}

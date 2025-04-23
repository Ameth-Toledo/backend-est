package controllers

import (
	"estsoftware/src/comments/application"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetCommentByIdController struct {
	useCase *application.GetCommentById
}

func NewGetCommentByIdController(useCase *application.GetCommentById) *GetCommentByIdController {
	return &GetCommentByIdController{useCase: useCase}
}

func (cc *GetCommentByIdController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	comment, err := cc.useCase.Execute(id)
	if err != nil {
		fmt.Println("ERROR real al obtener comentario:", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el comentario"})
		return
	}

	if comment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

package controllers

import (
	"estsoftware/src/comments/application"
	"estsoftware/src/comments/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateCommentController struct {
	useCase *application.UpdateComment
}

func NewUpdateCommentController(useCase *application.UpdateComment) *UpdateCommentController {
	return &UpdateCommentController{useCase: useCase}
}

func (cc *UpdateCommentController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var comment entities.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	comment.ID = int32(id)

	updated, err := cc.useCase.Execute(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el comentario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": updated})
}

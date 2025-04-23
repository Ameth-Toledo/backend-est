package controllers

import (
	"estsoftware/src/comments/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteCommentController struct {
	useCase *application.DeleteComment
}

func NewDeleteCommentController(useCase *application.DeleteComment) *DeleteCommentController {
	return &DeleteCommentController{useCase: useCase}
}

func (cc *DeleteCommentController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := cc.useCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el comentario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comentario eliminado correctamente"})
}

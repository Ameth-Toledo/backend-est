package controllers

import (
	"estsoftware/src/comments/application"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllCommentsController struct {
	useCase *application.GetAllComments
}

func NewGetAllCommentsController(useCase *application.GetAllComments) *GetAllCommentsController {
	return &GetAllCommentsController{useCase: useCase}
}

func (cc *GetAllCommentsController) Execute(c *gin.Context) {
	comments, err := cc.useCase.Execute()
	if err != nil {
		fmt.Printf("Error al obtener los comentarios: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los comentarios"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

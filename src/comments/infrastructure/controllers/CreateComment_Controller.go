package controllers

import (
	"estsoftware/src/comments/application"
	"estsoftware/src/comments/domain/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateCommentController struct {
	useCase *application.CreateComment
}

func NewCreateCommentController(useCase *application.CreateComment) *CreateCommentController {
	return &CreateCommentController{useCase: useCase}
}

func (cc *CreateCommentController) Execute(context *gin.Context) {
	var comment entities.Comment
	if err := context.ShouldBindJSON(&comment); err != nil {
		context.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	if comment.Fecha.IsZero() {
		comment.Fecha = time.Now()
	}

	createdComment, err := cc.useCase.Execute(comment)
	if err != nil {
		context.JSON(500, gin.H{"error": fmt.Sprintf("No se pudo crear el comentario: %v", err)})
		return
	}

	context.JSON(201, gin.H{"comment": createdComment})
}

package controllers

import (
	"estsoftware/src/comments/application"
	"estsoftware/src/comments/domain/entities"
	"github.com/gin-gonic/gin"
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

	createdComment, err := cc.useCase.Execute(comment)
	if err != nil {
		context.JSON(500, gin.H{"error": "No se pudo crear el comentario"})
		return
	}

	context.JSON(201, gin.H{"comment": createdComment})
}

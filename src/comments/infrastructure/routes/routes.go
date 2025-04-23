package routes

import (
	"estsoftware/src/comments/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureCommentsRoutes(r *gin.Engine,
	createCommentController *controllers.CreateCommentController,
	getAllCommentsController *controllers.GetAllCommentsController,
	getByIdCommentController *controllers.GetCommentByIdController,
	updateCommentController *controllers.UpdateCommentController,
	deleteCommentController *controllers.DeleteCommentController,
) {
	r.POST("/comments", createCommentController.Execute)
	r.GET("/comments", getAllCommentsController.Execute)
	r.GET("/comments/:id", getByIdCommentController.Execute)
	r.PUT("/comments/:id", updateCommentController.Execute)
	r.DELETE("/comments/:id", deleteCommentController.Execute)
}

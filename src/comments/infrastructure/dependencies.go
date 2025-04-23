package infrastructure

import (
	"estsoftware/src/comments/application"
	"estsoftware/src/comments/infrastructure/adapters"
	"estsoftware/src/comments/infrastructure/controllers"
	"estsoftware/src/core"
)

type DependenciesComments struct {
	CreateCommentController  *controllers.CreateCommentController
	GetAllCommentController  *controllers.GetAllCommentsController
	GetByIDCommentController *controllers.GetCommentByIdController
	UpdateCommentController  *controllers.UpdateCommentController
	DeleteCommentController  *controllers.DeleteCommentController
}

func InitComments() *DependenciesComments {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &DependenciesComments{
		CreateCommentController:  controllers.NewCreateCommentController(application.NewCreateComment(ps)),
		GetAllCommentController:  controllers.NewGetAllCommentsController(application.NewGetAllComments(ps)),
		GetByIDCommentController: controllers.NewGetCommentByIdController(application.NewGetCommentById(ps)),
		UpdateCommentController:  controllers.NewUpdateCommentController(application.NewUpdateComment(ps)),
		DeleteCommentController:  controllers.NewDeleteCommentController(application.NewDeleteComment(ps)),
	}
}

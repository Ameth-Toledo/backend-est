package infrastructure

import (
	"estsoftware/src/core"
	"estsoftware/src/pages/application"
	"estsoftware/src/pages/infrastructure/adapters"
	"estsoftware/src/pages/infrastructure/controllers"
)

type DependenciesPages struct {
	CreatePageController  *controllers.CreatePageController
	GetAllPagesController *controllers.GetAllPagesController
	GetByIdPageController *controllers.GetPageByIdController
	UpdatePageController  *controllers.UpdatePageController
	DeletePageController  *controllers.DeletePageController
}

func InitPages() *DependenciesPages {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &DependenciesPages{
		CreatePageController:  controllers.NewCreatePageController(application.NewCreatePage(ps)),
		GetAllPagesController: controllers.NewGetAllPagesController(application.NewGetAllPages(ps)),
		GetByIdPageController: controllers.NewGetPageByIdController(application.NewGetPageById(ps)),
		UpdatePageController:  controllers.NewUpdatePageController(application.NewUpdatePage(ps)),
		DeletePageController:  controllers.NewDeletePageController(application.NewDeletePage(ps)),
	}
}

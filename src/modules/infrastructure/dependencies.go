package infrastructure

import (
	"estsoftware/src/core"
	"estsoftware/src/modules/application"
	"estsoftware/src/modules/infrastructure/adapters"
	"estsoftware/src/modules/infrastructure/controllers"
)

type DependenciesModules struct {
	CreateModuleController  *controllers.CreateModuleController
	GetAllModulesController *controllers.GetAllModulesController
	GetModuleByIdController *controllers.GetModuleByIdController
	UpdateModuleController  *controllers.UpdateModuleController
	DeleteModuleController  *controllers.DeleteModuleController
}

func InitModules() *DependenciesModules {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &DependenciesModules{
		CreateModuleController:  controllers.NewCreateModuleController(application.NewCreateModule(ps)),
		GetAllModulesController: controllers.NewGetAllModulesController(application.NewGetAllModules(ps)),
		GetModuleByIdController: controllers.NewGetModuleByIdController(application.NewGetModuleById(ps)),
		UpdateModuleController:  controllers.NewUpdateModuleController(application.NewUpdateModule(ps)),
		DeleteModuleController:  controllers.NewDeleteModuleController(application.NewDeleteModule(ps)),
	}
}

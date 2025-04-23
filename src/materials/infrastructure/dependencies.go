package infrastructure

import (
	"estsoftware/src/core"
	"estsoftware/src/materials/application"
	"estsoftware/src/materials/infrastructure/adapters"
	"estsoftware/src/materials/infrastructure/controllers"
)

type MaterialDependencies struct {
	CreateMaterialController  *controllers.CreateMaterialController
	GetAllMaterialsController *controllers.GetAllMaterialsController
	GetMaterialByIdController *controllers.GetMaterialByIdController
	UpdateMaterialController  *controllers.UpdateMaterialController
	DeleteMaterialController  *controllers.DeleteMaterialController
}

func InitMaterials() *MaterialDependencies {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &MaterialDependencies{
		CreateMaterialController:  controllers.NewCreateMaterialController(application.NewCreateMaterial(ps)),
		GetAllMaterialsController: controllers.NewGetAllMaterialsController(application.NewGetAllMaterials(ps)),
		GetMaterialByIdController: controllers.NewGetMaterialByIdController(application.NewGetMaterialById(ps)),
		UpdateMaterialController:  controllers.NewUpdateMaterialController(application.NewUpdateMaterial(ps)),
		DeleteMaterialController:  controllers.NewDeleteMaterialController(application.NewDeleteMaterial(ps)),
	}
}

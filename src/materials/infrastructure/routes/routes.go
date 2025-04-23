package routes

import (
	"estsoftware/src/materials/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureMaterialRoutes(
	r *gin.Engine,
	createMaterialController *controllers.CreateMaterialController,
	getAllMaterialController *controllers.GetAllMaterialsController,
	getByIdMaterialController *controllers.GetMaterialByIdController,
	updateMaterialController *controllers.UpdateMaterialController,
	deleteMaterialController *controllers.DeleteMaterialController,
) {
	r.GET("/materials/:id", getByIdMaterialController.Execute)
	r.GET("/materials", getAllMaterialController.Execute)
	r.Static("/static/materials", "./materials")
	r.POST("/materials", createMaterialController.Execute)
	r.PUT("/materials", updateMaterialController.Execute)
	r.DELETE("/materials/:id", deleteMaterialController.Execute)
}

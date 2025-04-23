package routes

import (
	"estsoftware/src/modules/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureModulesRoutes(r *gin.Engine,
	createModuleController *controllers.CreateModuleController,
	getAllModulesController *controllers.GetAllModulesController,
	getModuleByIdController *controllers.GetModuleByIdController,
	updateModuleController *controllers.UpdateModuleController,
	deleteModuleController *controllers.DeleteModuleController,
) {
	r.Static("/image_modulo", "./image_modulo")

	r.POST("/modules", createModuleController.Execute)
	r.GET("/modules", getAllModulesController.Execute)
	r.GET("/modules/:id", getModuleByIdController.Execute)
	r.PUT("/modules/:id", updateModuleController.Execute)
	r.DELETE("/modules/:id", deleteModuleController.Execute)
}

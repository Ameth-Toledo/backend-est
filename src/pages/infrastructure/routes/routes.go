package routes

import (
	"estsoftware/src/pages/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigurePagesRoutes(r *gin.Engine,
	createPageController *controllers.CreatePageController,
	getAllPagesController *controllers.GetAllPagesController,
	getPageByIdController *controllers.GetPageByIdController,
	updatePageController *controllers.UpdatePageController,
	deletePageController *controllers.DeletePageController,
) {
	r.POST("/pages", createPageController.Execute)
	r.GET("/pages", getAllPagesController.Execute)
	r.GET("/pages/:id", getPageByIdController.Execute)
	r.PUT("/pages/:id", updatePageController.Execute)
	r.DELETE("/pages/:id", deletePageController.Execute)
}

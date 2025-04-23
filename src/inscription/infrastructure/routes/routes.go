package routes

import (
	"estsoftware/src/inscription/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureInscriptionRoutes(r *gin.Engine,
	createInscriptionController *controllers.CreateInscriptionController,
	getAllInscriptionsController *controllers.GetAllInscriptionsController,
	getInscriptionByIDController *controllers.GetInscriptionByIDController,
	deleteInscriptionController *controllers.DeleteInscriptionController,
	updateInscriptionController *controllers.UpdateInscriptionController,
) {
	r.POST("/inscriptions", createInscriptionController.Execute)
	r.GET("/inscriptions", getAllInscriptionsController.Execute)
	r.GET("/inscriptions/:id", getInscriptionByIDController.Execute)
	r.DELETE("/inscriptions/:id", deleteInscriptionController.Execute)
	r.PUT("/inscriptions/:id", updateInscriptionController.Execute)
}

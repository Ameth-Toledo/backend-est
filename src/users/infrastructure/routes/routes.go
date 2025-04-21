package routes

import (
	"estsoftware/src/core/security"
	"estsoftware/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureUserRoutes(r *gin.Engine,
	createClientController *controllers.CreateUserController,
	viewClientController *controllers.ViewUserController,
	editClientController *controllers.EditUserController,
	deleteClientController *controllers.DeleteUserController,
	viewByIdClientController *controllers.ViewUserByIdController,
	loginController *controllers.AuthController,
) {
	r.POST("/users", createClientController.Execute)
	r.GET("/users", security.JWTMiddleware(), viewClientController.Execute)
	r.POST("/login", loginController.Execute)
	r.GET("/users/:id", security.JWTMiddleware(), viewByIdClientController.Execute)
	r.PUT("/users/:id", security.JWTMiddleware(), editClientController.Execute)
	r.DELETE("/users/:id", security.JWTMiddleware(), deleteClientController.Execute)
}

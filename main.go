package main

import (
	userInfra "estsoftware/src/users/infrastructure"
	userRoutes "estsoftware/src/users/infrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.Static("uploads", "./uploads")

	userDependencies := userInfra.InitUsers()

	userRoutes.ConfigureUserRoutes(r,
		userDependencies.CreateUserController,
		userDependencies.ViewUserController,
		userDependencies.EditUserController,
		userDependencies.DeleteUserController,
		userDependencies.ViewUserByIdController,
		userDependencies.AuthController,
	)

	r.Run(":8080")
}

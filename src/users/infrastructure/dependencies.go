package infrastructure

import (
	"estsoftware/src/core"
	"estsoftware/src/users/application"
	"estsoftware/src/users/infrastructure/adapters"
	"estsoftware/src/users/infrastructure/controllers"
)

type DependenciesUsers struct {
	CreateUserController   *controllers.CreateUserController
	ViewUserController     *controllers.ViewUserController
	ViewUserByIdController *controllers.ViewUserByIdController
	EditUserController     *controllers.EditUserController
	DeleteUserController   *controllers.DeleteUserController
	AuthController         *controllers.AuthController
}

func InitUsers() *DependenciesUsers {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &DependenciesUsers{
		CreateUserController:   controllers.NewCreateUserController(*application.NewCreateUser(ps)),
		ViewUserController:     controllers.NewViewUserController(*application.NewListUser(ps)),
		ViewUserByIdController: controllers.NewViewUserByIdController(*application.NewUserById(ps)),
		EditUserController:     controllers.NewEditUserController(*application.NewEditUser(ps)),
		DeleteUserController:   controllers.NewDeleteUserController(*application.NewDeleteUser(ps)),
		AuthController:         controllers.NewAuthController(application.NewAuthService(ps)),
	}
}

package main

import (
	commentInfra "estsoftware/src/comments/infrastructure"
	commentRoutes "estsoftware/src/comments/infrastructure/routes"
	courseInfra "estsoftware/src/course/infrastructure"
	courseRoutes "estsoftware/src/course/infrastructure/routes"
	inscriptionInfra "estsoftware/src/inscription/infrastructure"
	inscriptionRoutes "estsoftware/src/inscription/infrastructure/routes"
	materialInfra "estsoftware/src/materials/infrastructure"
	materialRoutes "estsoftware/src/materials/infrastructure/routes"
	moduleInfra "estsoftware/src/modules/infrastructure"
	moduleRoutes "estsoftware/src/modules/infrastructure/routes"
	pageInfra "estsoftware/src/pages/infrastructure"
	pageRoutes "estsoftware/src/pages/infrastructure/routes"
	userInfra "estsoftware/src/users/infrastructure"
	userRoutes "estsoftware/src/users/infrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.StaticFS("/img", gin.Dir("./static/img", false))
	r.StaticFS("/uploads", gin.Dir("./uploads", false))

	userDependencies := userInfra.InitUsers()

	userRoutes.ConfigureUserRoutes(r,
		userDependencies.CreateUserController,
		userDependencies.ViewUserController,
		userDependencies.EditUserController,
		userDependencies.DeleteUserController,
		userDependencies.ViewUserByIdController,
		userDependencies.AuthController,
	)

	courseDependencies := courseInfra.InitCourses()
	courseRoutes.ConfigureCoursesRoutes(r,
		courseDependencies.CreateCourseController,
		courseDependencies.GetCoursesController,
		courseDependencies.GetCourseByIdController,
		courseDependencies.GetCourseByNameController,
		courseDependencies.UpdateCourseController,
		courseDependencies.DeleteCourseController,
	)

	moduleDependencies := moduleInfra.InitModules()
	moduleRoutes.ConfigureModulesRoutes(r,
		moduleDependencies.CreateModuleController,
		moduleDependencies.GetAllModulesController,
		moduleDependencies.GetModuleByIdController,
		moduleDependencies.UpdateModuleController,
		moduleDependencies.DeleteModuleController,
	)

	pageDependencies := pageInfra.InitPages()
	pageRoutes.ConfigurePagesRoutes(r,
		pageDependencies.CreatePageController,
		pageDependencies.GetAllPagesController,
		pageDependencies.GetByIdPageController,
		pageDependencies.UpdatePageController,
		pageDependencies.DeletePageController,
	)

	materialDependencies := materialInfra.InitMaterials()
	materialRoutes.ConfigureMaterialRoutes(r,
		materialDependencies.CreateMaterialController,
		materialDependencies.GetAllMaterialsController,
		materialDependencies.GetMaterialByIdController,
		materialDependencies.UpdateMaterialController,
		materialDependencies.DeleteMaterialController,
	)

	inscriptionDependencies := inscriptionInfra.InitInscription()
	inscriptionRoutes.ConfigureInscriptionRoutes(r,
		inscriptionDependencies.CreateInscription,
		inscriptionDependencies.GetAllInscriptions,
		inscriptionDependencies.GetInscriptionByID,
		inscriptionDependencies.DeleteInscription,
		inscriptionDependencies.UpdateInscription,
	)

	commentsDependencies := commentInfra.InitComments()
	commentRoutes.ConfigureCommentsRoutes(r,
		commentsDependencies.CreateCommentController,
		commentsDependencies.GetAllCommentController,
		commentsDependencies.GetByIDCommentController,
		commentsDependencies.UpdateCommentController,
		commentsDependencies.DeleteCommentController,
	)

	r.Run(":8080")
}

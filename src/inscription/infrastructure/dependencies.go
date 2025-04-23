package infrastructure

import (
	"estsoftware/src/core"
	"estsoftware/src/inscription/application"
	"estsoftware/src/inscription/infrastructure/adapters"
	"estsoftware/src/inscription/infrastructure/controllers"
)

type DependenciesInscription struct {
	CreateInscription  *controllers.CreateInscriptionController
	GetAllInscriptions *controllers.GetAllInscriptionsController
	GetInscriptionByID *controllers.GetInscriptionByIDController
	DeleteInscription  *controllers.DeleteInscriptionController
	UpdateInscription  *controllers.UpdateInscriptionController
}

func InitInscription() *DependenciesInscription {
	conn := core.GetDBPool()
	ps := adapters.NewMySQL(conn.DB)

	return &DependenciesInscription{
		CreateInscription:  controllers.NewCreateInscripcionController(application.NewCreateInscription(ps)),
		GetAllInscriptions: controllers.NewGetAllInscriptionsController(application.NewGetAllInscriptions(ps)),
		GetInscriptionByID: controllers.NewGetInscriptionByIDController(application.NewGetInscriptionByID(ps)),
		DeleteInscription:  controllers.NewDeleteInscriptionController(application.NewDeleteInscription(ps)),
		UpdateInscription:  controllers.NewUpdateInscriptionController(application.NewUpdateInscription(ps)),
	}
}

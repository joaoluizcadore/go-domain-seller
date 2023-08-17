package application

import (
	"database/sql"

	"github.com/joaoluizcadore/domain-seller/infra/database"
	"github.com/joaoluizcadore/domain-seller/internal/service"
)

var (
	instance *Application
)

type Application struct {
	VisitService        *service.VisitCountService
	NotificationService *service.NotificationService
}

func InitializeApplication(db *sql.DB) {
	database.CreateSqlDDL(db)
	visitRepository := database.NewVisitCountSqlRepository(db)
	visitService := service.NewVisitCountService(visitRepository)
	notificationService := service.NewNotificationService()

	instance = &Application{
		VisitService:        visitService,
		NotificationService: notificationService,
	}
}

func GetApp() *Application {
	if instance == nil {
		panic("Application not initialized")
	}
	return instance
}

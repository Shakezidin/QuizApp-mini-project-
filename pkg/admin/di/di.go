package di

import (
	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/admin/handler"
	"github.com/Shakezidin/pkg/admin/repository"
	"github.com/Shakezidin/pkg/admin/routes"
	"github.com/Shakezidin/pkg/admin/service"
	"github.com/Shakezidin/pkg/db"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, cnfg *config.Config) {

	// Connect to database
	dbConn := db.Database(cnfg)

	// Create repository
	adminRepo := repository.NewAdminRepo(dbConn)

	// Create user service
	adminService := service.NewAdminSVC(adminRepo, cnfg)

	// Create user handler
	adminHandler := handler.NewUserHandler(adminService)

	routes.AdminRoutes(&adminHandler, r)
}

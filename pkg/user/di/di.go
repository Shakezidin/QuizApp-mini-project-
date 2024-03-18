package di

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/db"
	"github.com/Shakezidin/pkg/user/handler"
	"github.com/Shakezidin/pkg/user/repository"
	"github.com/Shakezidin/pkg/user/routes"
	"github.com/Shakezidin/pkg/user/service"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, cnfg *config.Config) {
	// Connect to Redis
	redis, err := config.ConnectToRedis(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	// Connect to database
	dbConn := db.Database(cnfg)

	// Create repository
	adminRepo := repository.NewAdminRepo(dbConn)

	// Create user service
	adminService := service.NewAdminSVC(adminRepo, redis, cnfg)

	// Create user handler
	adminHandler := handler.NewUserHandler(adminService)

	routes.AdminRoutes(adminHandler, r)
}

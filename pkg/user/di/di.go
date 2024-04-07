package di

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/db"
	"github.com/Shakezidin/pkg/user/handler"
	"github.com/Shakezidin/pkg/user/repository"
	"github.com/Shakezidin/pkg/user/routes"
	"github.com/Shakezidin/pkg/user/service"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, cnfg config.Config) {

	// Connect to database
	dbConn := db.Database(cnfg)

	redis, err := config.ConnectToRedis(cnfg)
	if err != nil {
		log.Println("error while connecting redis")
		fmt.Println(err)
		return
	}

	smtp := config.NewSMTP(redis)
	// Create repository
	adminRepo := repository.NewAdminRepo(dbConn)

	// Create user service
	adminService := service.NewAdminSVC(adminRepo, cnfg, &smtp, redis)

	// Create user handler
	adminHandler := handler.NewUserHandler(adminService)

	routes.AdminRoutes(&adminHandler, r)
}

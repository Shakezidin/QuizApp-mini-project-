package routes

import (
	"github.com/Shakezidin/pkg/admin/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(adminhandler *handler.AdminHandler, r *gin.Engine) {
	apiVersion := r.Group("/api/v1")

	admin := apiVersion.Group("/admin")
	{
		admin.POST("/login",adminhandler.AdminLogin)
	}
}

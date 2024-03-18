package routes

import "github.com/gin-gonic/gin"

func AdminRoutes(r *gin.Engine) {
	apiVersion := r.Group("/api/v1")

	user := apiVersion.Group("/user")
	{

	}
}

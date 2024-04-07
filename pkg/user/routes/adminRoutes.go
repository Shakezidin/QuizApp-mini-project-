package routes

import (
	"fmt"

	"github.com/Shakezidin/pkg/user/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(userhandler *handler.AdminHandler, r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		fmt.Println("Hello World!") // Test if this route works
	})

	apiVersion := r.Group("/api/v1")
	user := apiVersion.Group("/user")
	{
		user.POST("/signup", userhandler.AdminSignup)
		user.POST("/signup/verify", userhandler.SignupVerify)
		user.POST("/login", userhandler.AdminLogin)

		user.GET("/quiz/languages", userhandler.GetQuizLanguages)
		user.GET("/quiz/start/:language", userhandler.StartQuiz) // Endpoint to generate quiz question
		user.POST("/quiz/check", userhandler.CheckAnswer)
	}
}

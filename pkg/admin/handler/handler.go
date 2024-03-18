package handler

import (
	"fmt"
	"net/http"

	"github.com/Shakezidin/pkg/admin/service"
	"github.com/Shakezidin/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminHandler struct {
	AdminSVC service.AdminService
}

func (a *AdminHandler) AdminLogin(c *gin.Context) {
	var login model.LoginCredentials

	if err := c.BindJSON(&login); err != nil {
		c.JSON(400, gin.H{
			"Error": "Binding Error",
		})
		return
	}

	validation := validator.New()

	if err := validation.Struct(&login); err != nil {
		errMsg := "Validation error"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  errMsg,
		})
		return
	}

	response, err := a.AdminSVC.AdminLoginSVC(login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": fmt.Sprintf("%s logged in successfully", login.Username),
		"Data":    response,
	})
}

func NewUserHandler(adminsvc service.AdminService) AdminHandler {
	return AdminHandler{
		AdminSVC: adminsvc,
	}
}

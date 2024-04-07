package handler

import (
	"fmt"
	"net/http"

	"github.com/Shakezidin/pkg/model"
	"github.com/Shakezidin/pkg/user/service"
	"github.com/Shakezidin/utils"
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

func (a *AdminHandler) AdminSignup(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	// Validate struct
	validate := validator.New()
	validate.RegisterValidation("emailcst", utils.EmailValidation)
	validate.RegisterValidation("phone", utils.PhoneNumberValidation)
	validate.RegisterValidation("alphaspace", utils.AlphaSpace)
	err := validate.Struct(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": fmt.Sprintf("Error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	response, err := a.AdminSVC.UserSignupSVC(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status": http.StatusAccepted,
		"Data":   response,
	})
}

func (a *AdminHandler) SignupVerify(c *gin.Context) {
	var OtpVerify model.OtpVerify
	if err := c.BindJSON(&OtpVerify); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	response, err := a.AdminSVC.UserSignupVerify(OtpVerify)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": "OTP verified, user creation successful.",
		"Data":    response,
	})
}

func NewUserHandler(adminsvc service.AdminService) AdminHandler {
	return AdminHandler{
		AdminSVC: adminsvc,
	}
}

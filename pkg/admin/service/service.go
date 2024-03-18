package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/admin/repository"
	"github.com/Shakezidin/pkg/model"
	"github.com/Shakezidin/utils"
	"gorm.io/gorm"
)

type AdminService struct {
	AdminRepo repository.AdminRepository
	cnfg      *config.Config
}

func (a *AdminService) AdminLoginSVC(admin model.LoginCredentials) (string, error) {
	// Find user by email
	user, err := a.AdminRepo.FindAdminRepo(admin.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", errors.New("unable to login")
	}

	// Check password
	check := utils.CheckPasswordMatch([]byte(user.Password), admin.Password)
	if !check {
		return "", fmt.Errorf("password mismatch for admin %v", admin.Username)
	}

	// Generate token
	userIdstr := strconv.Itoa(int(user.ID))
	token, _, err := utils.GenerateToken(admin.Username, "admin", userIdstr, a.cnfg.SECRETKEY)
	if err != nil {
		return "", errors.New("error while generating token")
	}

	return token, nil
}

func NewAdminSVC(adminrepo repository.AdminRepository, cnfg *config.Config) AdminService {
	return AdminService{
		AdminRepo: adminrepo,
		cnfg:      cnfg,
	}
}

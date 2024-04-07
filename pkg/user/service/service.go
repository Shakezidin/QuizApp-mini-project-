package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/model"
	"github.com/Shakezidin/pkg/user/repository"
	"github.com/Shakezidin/utils"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AdminService struct {
	AdminRepo   repository.AdminRepository
	cnfg        config.Config
	Smtp        *config.Smtp
	RedisClient *redis.Client
}

func (a *AdminService) AdminLoginSVC(admin model.LoginCredentials) (*model.Responce, error) {
	// Find user by email
	user, err := a.AdminRepo.FindAdminRepo(admin.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("unable to login")
	}

	// Check password
	check := utils.CheckPasswordMatch([]byte(user.Password), admin.Password)
	if !check {
		return nil, fmt.Errorf("password mismatch for admin %v", admin.Username)
	}

	// Generate token
	userIdstr := strconv.Itoa(int(user.ID))
	token, _, err := utils.GenerateToken(admin.Username, "admin", userIdstr, a.cnfg.SECRETKEY)
	if err != nil {
		return nil, errors.New("error while generating token")
	}

	return &model.Responce{
		Token: token,
	}, nil
}

func (a *AdminService) UserSignupSVC(user model.User) (*model.Responce, error) {
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("error while hashing password")
	}

	user.Password = string(hashPassword)

	Otp := a.Smtp.GetOTP(user.Username, user.Email)

	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Inserting the OTP into Redis
	err = a.RedisClient.Set(context.Background(), "signUpOTP"+user.Email, Otp, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	// Inserting the data into Redis
	err = a.RedisClient.Set(context.Background(), "userData"+user.Email, jsonData, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	return &model.Responce{
		Status: "user signup initiated",
	}, nil
}

func (a *AdminService) UserSignupVerify(otpCred model.OtpVerify) (*model.Responce, error) {
	if a.Smtp.VerifyOTP("signUpOTP"+otpCred.Email, otpCred.OTP) {
		var userData model.User
		superKey := "userData" + otpCred.Email
		jsonData, err := a.RedisClient.Get(context.Background(), superKey).Result()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(jsonData), &userData)
		if err != nil {
			return nil, err
		}

		// Create user and save transaction history and wallet balance
		results, err := a.AdminRepo.CreateUser(userData)
		if err != nil {
			return nil, err
		}
		return &model.Responce{
			Id:     int(results.ID),
			Status: "User signup success",
		}, nil
	}

	return &model.Responce{
		Status: "User signup failed",
	}, errors.New("user signup failed")
}

func NewAdminSVC(adminrepo repository.AdminRepository, cnfg config.Config, smtp *config.Smtp, redis *redis.Client) AdminService {
	return AdminService{
		AdminRepo:   adminrepo,
		cnfg:        cnfg,
		Smtp:        smtp,
		RedisClient: redis,
	}
}

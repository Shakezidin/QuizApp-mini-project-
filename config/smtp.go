package config

import (
	"context"
	"crypto/rand"
	"math/big"
	"net/smtp"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Smtp struct {
	ReddisClient *redis.Client
}

// GetOTP generates and sends a one-time password (OTP) via email.
func (s *Smtp) GetOTP(name, email string) string {
	otp, err := getRandNum()
	if err != nil {
		panic(err)
	}
	msg := "Subject: WebPortal OTP\nHey " + name + "Your OTP is " + otp
	s.sendEmail(name, msg, email)
	return otp
}

// getRandNum generates a random OTP.
func getRandNum() (string, error) {
	otp, err := rand.Int(rand.Reader, big.NewInt(8999))
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(otp.Int64()+1000, 10), nil
}

// sendEmail sends an email with the OTP.
func (s *Smtp) sendEmail(name, msg, email string) {
	SMTPemail := os.Getenv("EMAIL")
	SMTPpass := os.Getenv("PASSWORD")
	auth := smtp.PlainAuth("", SMTPemail, SMTPpass, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, SMTPemail, []string{email}, []byte(msg))
	if err != nil {
		panic(err)
	}
}

// VerifyOTP verifies if the provided OTP matches the stored OTP in Redis.
func (s *Smtp) VerifyOTP(superkey, otpInput string) bool {
	// OTP verification in Redis
	otp, err := s.ReddisClient.Get(context.Background(), superkey).Result()
	if err != nil {
		return false
	}

	if otp == otpInput {
		err := s.ReddisClient.Del(context.Background(), superkey).Err()
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func NewSMTP(redis *redis.Client) Smtp {
	return Smtp{
		ReddisClient: redis,
	}
}

package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Id    string
	Email string
	Role  string
	Token string
	jwt.StandardClaims
}

func GenerateToken(email, role, id string, cfg string) (string, string, error) {
	expireTime := time.Now().Add(time.Hour * 4).Unix()
	claims := &Claims{
		Id:    id,
		Email: email,
		Role:  role,
		Token: "Access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg))
	if err != nil {
		log.Printf("unable to generate jwt token for user %v, err: %v", email, err.Error())
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Id:    id,
		Token: "Refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 240).Unix(),
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	})
	rt, err := refreshToken.SignedString([]byte(cfg))
	if err != nil {
		return "", "", err
	}

	return signedToken, rt, nil
}

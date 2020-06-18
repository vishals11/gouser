package model

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vishals11/gouser/config"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(user *User) error {
	var err error
	cfg := config.Get()
	key := []byte(cfg.JwtSigningKey)

	claims := Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	user.Token, err = token.SignedString(key)
	if err != nil {
		log.Printf("Error Signing string:%v", err)
	}

	return err
}

func ValidateToken(tokenString string) (*User, error) {
	cfg := config.Get()
	key := []byte(cfg.JwtSigningKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Printf("Error parsing the JWT Token: %s", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Error reading claims: %v", err)
		return nil, err
	}

	userID := int(claims["user_id"].(float64))
	user, err := GetUserFromID(userID)
	if err != nil {
		return nil, err
	}

	if user.Token != tokenString {
		return nil, fmt.Errorf("Token is Invalid. Login again")
	}

	return user, err
}

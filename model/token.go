package model

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vishals11/gouser/config"
)

func GenerateToken(user *User) error {
	var err error
	cfg := config.Get()
	key := []byte(cfg.JwtSigningKey)

	claims := &jwt.StandardClaims{
		Id:        strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	user.Token, err = token.SignedString(key)
	return err
}

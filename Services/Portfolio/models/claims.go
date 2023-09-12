package models

import (
	app_config "libery_labs_portfolio/Config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

type ChatClaims struct {
	ChatID   string `json:"chat_id"`
	ClientIP string `json:"client_ip"`
	jwt.StandardClaims
}

func CreateChatClaims(client_ip string) (error, string) {
	chat_id := uuid.New().String()

	claims_verifier := argon2.IDKey([]byte(client_ip+app_config.JWT_SECRET), []byte(client_ip), 1, 64*1024, 4, 32)

	claims := ChatClaims{
		ChatID:   chat_id,
		ClientIP: client_ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(), // 30 minutes
			Issuer:    "Portfolio",
		},
	}

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwt_token.SignedString(claims_verifier)
	return err, token
}

func VerifyChatClaims(token string, client_ip string) (error, *ChatClaims) {
	claims := new(ChatClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return argon2.IDKey([]byte(client_ip+app_config.JWT_SECRET), []byte(client_ip), 1, 64*1024, 4, 32), nil
	})

	if err != nil {
		return err, nil
	}

	return nil, claims
}

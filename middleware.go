package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// getTokenFromHeader 함수는 쿠키에 저장된 Token, SignKey 정보를 읽어와 토큰키를 반환하는 함수이다.
func getTokenFromHeader(context *gin.Context) (Token, error) {
	var tk Token

	sessionToken, err := context.Cookie("SessionToken")
	if err != nil {
		return tk, err
	}

	sessionSignKey, err := context.Cookie("SessionSignKey")
	if err != nil {
		return tk, err
	}

	// Signkey로 Token 정보를 연다.
	token, err := jwt.ParseWithClaims(sessionToken, &tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(sessionSignKey), nil
	})
	if err != nil {
		return tk, err
	}

	if !token.Valid {
		return tk, errors.New("Token key is not valid")
	}

	return tk, nil
}
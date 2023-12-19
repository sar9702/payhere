package main

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// getTokenFromCookie 함수는 쿠키에 저장된 Token, SignKey 정보를 읽어와 토큰키를 반환하는 함수이다.
func getTokenFromCookie(context *gin.Context) (Token, error) {
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

// checkTokenFromHeader 함수는 Rest API 사용 시 헤더에 저장된 토큰이 유효한지 확인하는 함수이다.
func checkTokenFromHeader(context *gin.Context) error {
	headerValue := context.Request.Header["Authorization"]
	if len(headerValue) != 1 {
		return errors.New("Authorization Failed")
	}

	//header에서 token을 가져온다.
	auth := strings.SplitN(headerValue[0], " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return errors.New("Authorization Failed")
	}
	token := auth[1]

	_, err := userByToken(token)
	if err != nil {
		return err
	}
	
	return nil
}
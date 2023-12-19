package main

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Item struct {
	ID string `form:"id"` // ID
	
	Category       string `form:"category"`        // 카테고리
	Name           string `form:"name"`            // 이름
	Price          string `form:"price"`           // 가격
	Cost           string `form:"cost"`            // 원가
	Description    string `form:"description"`     // 설명
	Barcode        string `form:"barcode"`         // 바코드
	ExpirationDate string `form:"expirationDate"` // 유통기한
	Size           string `form:"size"`            // 사이즈 small or large
}

// Token 자료구조. JWT 방식을 사용한다. restAPI 사용시 보안체크를 위해 http 헤더에 들어간다.
type Token struct {
	ID          string `json:"id" bson:"id"`                   // 사용자 ID
	jwt.StandardClaims
}

// User 자료구조
type User struct {
	ID string // 사용자 ID(핸드폰번호)
	Password string // 암호화된 비밀번호
	Token string // JWT 토큰
	SignKey string // JWT 토큰을 만들 때 사용하는 SignKey
}

// CreateToken 메소드는 토큰을 생성합니다.
func (u *User) CreateToken() error {
	if u.ID == "" {
		return errors.New("ID is an empty string")
	}
	if u.Password == "" {
		return errors.New("Password is an empty string")
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{
		ID:          u.ID,
	})
	signKey, err := encrypt(u.Password)
	if err != nil {
		return err
	}
	u.SignKey = signKey
	tokenString, err := token.SignedString([]byte(signKey))
	if err != nil {
		return err
	}
	u.Token = tokenString
	return nil
}
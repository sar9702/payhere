package main

import (
	"golang.org/x/crypto/bcrypt"
)

// encrypt 함수는 문자를 받아서 해쉬문자로 변환한다.
func encrypt(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
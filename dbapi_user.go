package main

import (
	"database/sql"
	"errors"
)

// addUser 함수는 DB에 사용자를 등록하는 함수이다.
func addUser(user User) error {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	defer db.Close()

	// 동일한 ID를 가진 사용자가 있는지 확인한다.
	var exists bool
	err = db.QueryRow("SELECT COUNT(*) FROM User WHERE ID=?", user.ID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("이미 등록된 전화번호입니다")
	}

	_, err = db.Exec("INSERT INTO User VALUES (?, ?, ?, ?)", user.ID, user.Password, user.Token, user.SignKey)
	if err != nil {
		return err
	}

	return nil
}

// userByID 함수는 DB에서 아이디로 사용자를 찾아 반환하는 함수이다.
func userByID(id string) (User, error) {
	var user User

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return user, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM User WHERE ID=?", id).Scan(&user.ID, &user.Password, &user.Token, &user.SignKey)
	if err != nil {
		return user, err
	}

	return user, nil
}

// userByToken 함수는 DB에서 토큰키로 사용자를 찾아 반환하는 함수이다.
func userByToken(token string) (User, error) {
	var user User

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return user, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM User WHERE Token=?", token).Scan(&user.ID, &user.Password, &user.Token, &user.SignKey)
	if err != nil {
		return user, err
	}

	return user, nil
}
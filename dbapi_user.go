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
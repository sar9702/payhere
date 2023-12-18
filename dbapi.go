package main

import (
	"database/sql"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

var dns = "root:@tcp(localhost:3306)/payhere"

// items 함수는 DB에서 모든 아이템 리스트를 가져오는 함수이다.
func items() ([]Item, error) {
	var items []Item

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return items, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Item")
	if err != nil {
		return items, err
	}
	
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Category, &item.Name, &item.Price, &item.Cost, &item.Description, &item.Barcode, &item.ExpirationDate, &item.Size)
		if err != nil {
			return []Item{}, err
		}
		items = append(items, item)
	}

	return items, nil
}

// itemByID 함수는 DB에서 ID로 아이템을 찾아 반환하는 함수이다.
func itemByID(id string) (Item, error) {
	var item Item

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return item, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM Item WHERE ID=?", id).Scan(&item.ID, &item.Category, &item.Name, &item.Price, &item.Cost, &item.Description, &item.Barcode, &item.ExpirationDate, &item.Size)
	if err != nil {
		return item, err
	}

	return item, nil
}

// addItem 함수는 DB에 아이템을 추가하는 함수이다.
func addItem(item Item) error {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	defer db.Close()

	// 아이템 ID 생성
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO Item VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", uuid.String(), item.Category, item.Name, item.Price, item.Cost, item.Description, item.Barcode, item.ExpirationDate, item.Size)
	if err != nil {
		return err
	}

	return nil
}
package main

import (
	"database/sql"

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
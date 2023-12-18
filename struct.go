package main

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
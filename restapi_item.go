package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleAPIItemRegister 함수는 사용자가 입력한 데이터를 토대로 아이템을 등록하는 함수이다.
func handleAPIItemRegister(context *gin.Context) {
	// 헤더에 저장된 토큰키가 유효한지 확인한다.
	err := checkTokenFromHeader(context)
	if err != nil {
		var code int
		if err.Error() == "Authorization Failed" {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}

		jsonData := genResponseJson(code, err.Error(), nil)
		context.JSON(code, jsonData)
		return
	}

	var item Item

	// POST request 데이터를 item에 넣는다.
	err = context.Bind(&item)
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	err = addItem(item)
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	jsonData := genResponseJson(http.StatusOK, "ok", nil)
	context.JSON(http.StatusOK, jsonData)
}

// handleAPIItems 함수는 DB에 저장된 모든 아이템 리스트를 반환하는 API이다.
func handleAPIItems(context *gin.Context) {
	// 헤더에 저장된 토큰키가 유효한지 확인한다.
	err := checkTokenFromHeader(context)
	if err != nil {
		var code int
		if err.Error() == "Authorization Failed" {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}

		jsonData := genResponseJson(code, err.Error(), nil)
		context.JSON(code, jsonData)
		return
	}

	searchWord := context.Query("searchword")
	if searchWord == "null" {
		searchWord = ""
	}

	var items []Item
	if searchWord == "" {
		items, err = allItems()
		if err != nil {
			jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
			context.JSON(http.StatusInternalServerError, jsonData)
			return
		}
	} else {
		items, err = searchItem(searchWord)
		if err != nil {
			jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
			context.JSON(http.StatusInternalServerError, jsonData)
			return
		}
	}

	jsonData := genResponseJson(http.StatusOK, "ok", items)
	context.JSON(http.StatusOK, jsonData)
}

// handleAPIItemByID 함수는 ID가 일치하는 아이템을 찾아 반환하는 API이다.
func handleAPIItemByID(context *gin.Context) {
	// 헤더에 저장된 토큰키가 유효한지 확인한다.
	err := checkTokenFromHeader(context)
	if err != nil {
		var code int
		if err.Error() == "Authorization Failed" {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}

		jsonData := genResponseJson(code, err.Error(), nil)
		context.JSON(code, jsonData)
		return
	}

	id := context.Param("id")

	item, err := itemByID(id)
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	jsonData := genResponseJson(http.StatusOK, "ok", item)
	context.JSON(http.StatusOK, jsonData)
}

// handleAPIItemUpdateByID 함수는 아이템을 수정하는 API이다.
func handleAPIItemUpdateByID(context *gin.Context) {
	// 헤더에 저장된 토큰키가 유효한지 확인한다.
	err := checkTokenFromHeader(context)
	if err != nil {
		var code int
		if err.Error() == "Authorization Failed" {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}

		jsonData := genResponseJson(code, err.Error(), nil)
		context.JSON(code, jsonData)
		return
	}

	var item Item
	item.ID = context.Param("id")

	// PUT request 데이터를 item에 넣는다.
	err = context.Bind(&item)
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	err = updateItem(item)
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	jsonData := genResponseJson(http.StatusOK, "ok", item)
	context.JSON(http.StatusOK, jsonData)
}

// handleAPIItemDeleteByID 함수는 ID가 일치하는 아이템을 삭제하는 API이다.
func handleAPIItemDeleteByID(context *gin.Context) {
	// 헤더에 저장된 토큰키가 유효한지 확인한다.
	err := checkTokenFromHeader(context)
	if err != nil {
		var code int
		if err.Error() == "Authorization Failed" {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}

		jsonData := genResponseJson(code, err.Error(), nil)
		context.JSON(code, jsonData)
		return
	}

	id := context.Param("id")

	err = rmItem(id)
	if err != nil {
		jsonData := genResponseJson(http.StatusInternalServerError, err.Error(), nil)
		context.JSON(http.StatusInternalServerError, jsonData)
		return
	}

	jsonData := genResponseJson(http.StatusOK, "ok", nil)
	context.JSON(http.StatusOK, jsonData)
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleInit 함수는 메인 페이지를 띄우는 함수이다.
func handleInit(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "init", gin.H{
		"token": token,
	})
}

// handleItemDetail 함수는 아이템의 상세 페이지를 띄우는 함수이다.
func handleItemDetail(context *gin.Context) {
	// 로그인 정보가 유효한지 확인한다.
	token, err := getTokenFromCookie(context)
	if err != nil {
		context.Redirect(http.StatusSeeOther, "/signin")
		return
	}

	context.HTML(http.StatusOK, "detail", gin.H{
		"token": token,
	})
}

// // handleItemRegister 함수는 아이템 등록 페이지를 띄우는 함수이다.
// func handleItemRegister(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	token, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	context.HTML(http.StatusOK, "register", gin.H{
// 		"token": token,
// 	})
// }

// // handleItemRegisterSubmit 함수는 아이템 등록 과정을 처리하는 함수이다.
// func handleItemRegisterSubmit(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	_, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	var item Item

// 	// POST request 데이터를 item에 넣는다.
// 	err = context.Bind(&item)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	err = addItem(item)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	switch context.Request.Header.Get("Accept") {
// 	case "application/json":
// 		// Response with JSON
// 	default:
// 		context.Redirect(http.StatusSeeOther, "/item/register-success")
// 	}
// }

// // handleItemRegisterSuccess 함수는 아이템 등록 완료 페이지를 띄우는 함수이다.
// func handleItemRegisterSuccess(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	token, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	context.HTML(http.StatusOK, "register-success", gin.H{
// 		"token": token,
// 	})
// }

// // handleItemUpdate 함수는 아이템 수정 페이지를 띄우는 함수이다.
// func handleItemUpdate(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	token, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	id := context.Param("id")

// 	item, err := itemByID(id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	context.HTML(http.StatusOK, "edit", gin.H{
// 		"token": token,
// 		"product": item,
// 	})
// }

// // handleItemUpdateSubmit 함수는 아이템 수정 과정을 처리하는 함수이다.
// func handleItemUpdateSubmit(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	_, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	var item Item
// 	item.ID = context.Param("id")

// 	// POST request 데이터를 item에 넣는다.
// 	err = context.Bind(&item)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	err = updateItem(item)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	switch context.Request.Header.Get("Accept") {
// 	case "application/json":
// 		// Response with JSON
// 	default:
// 		context.Redirect(http.StatusSeeOther, "/item/update-success/" + item.ID)
// 	}
// }

// // handleItemUpdateSuccess 함수는 아이템 수정 완료 페이지를 띄우는 함수이다.
// func handleItemUpdateSuccess(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	token, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	id := context.Param("id")
	
// 	context.HTML(http.StatusOK, "edit-success", gin.H{
// 		"token": token,
// 		"id": id,
// 	})
// }

// // handleItemSearchSubmit 함수는 URL에 검색어 데이터를 포함하여 리다이렉트하는 함수이다.
// func handleItemSearchSubmit(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	_, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	searchWord := context.PostForm("searchWord")
// 	searchWord = strings.Trim(searchWord, " ")  // 앞뒤 공백 제거

// 	context.Redirect(http.StatusSeeOther, "/item/search?searchword=" + searchWord)
// }

// // handleItemSearch 함수는 아이템 검색 과정을 처리하는 함수이다.
// func handleItemSearch(context *gin.Context) {
// 	// 로그인 정보가 유효한지 확인한다.
// 	token, err := getTokenFromHeader(context)
// 	if err != nil {
// 		context.Redirect(http.StatusSeeOther, "/signin")
// 		return
// 	}

// 	searchWord := context.Query("searchword")

// 	items, err := searchItem(searchWord)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	switch context.Request.Header.Get("Accept") {
// 	case "application/json":
// 		// Response with JSON
// 	default:
// 		context.HTML(http.StatusOK, "init", gin.H{
// 			"token": token,
// 			"searchWord": searchWord,
// 			"products": items,
// 		})
// 	}
// }
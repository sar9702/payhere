package main

import (
	"strings"
)

// genResponseJson 함수는 API 호출 시 반환되어야 할 데이터를 반환하는 함수이다.
func genResponseJson(code int, msg string, data interface{}) map[string]interface{} {
	mapData := map[string]interface{}{
		"meta": map[string]interface{}{
			"code": code,
			"message": msg,
		},
		"data": data,
	}

	return mapData
}

// getChosung 함수는 문자열의 초성을 반환하는 함수이다.
func getChosung(name string) string {
	result := ""

	for _, s := range strings.Split(name, "") {
		if s >= "가" && s <= "깋" {
			result += "ㄱ"
		} else if s >= "까" && s <= "낗" {
			result += "ㄲ"
		} else if s >= "나" && s <= "닣" {
			result += "ㄴ"
		} else if s >= "다" && s <= "딯" {
			result += "ㄷ"
		} else if s >= "따" && s <= "띻" {
			result += "ㄸ"
		} else if s >= "라" && s <= "맇" {
			result += "ㄹ"
		} else if s >= "마" && s <= "밓" {
			result += "ㅁ"
		} else if s >= "바" && s <= "빟" {
			result += "ㅂ"
		} else if s >= "빠" && s <= "삫" {
			result += "ㅃ"
		} else if s >= "사" && s <= "싷" {
			result += "ㅅ"
		} else if s >= "싸" && s <= "앃" {
			result += "ㅆ"
		} else if s >= "아" && s <= "잏" {
			result += "ㅇ"
		} else if s >= "자" && s <= "짛" {
			result += "ㅈ"
		} else if s >= "짜" && s <= "찧" {
			result += "ㅉ"
		} else if s >= "차" && s <= "칳" {
			result += "ㅊ"
		} else if s >= "카" && s <= "킿" {
			result += "ㅋ"
		} else if s >= "타" && s <= "팋" {
			result += "ㅌ"
		} else if s >= "파" && s <= "핗" {
			result += "ㅍ"
		} else if s >= "하" && s <= "힣" {
			result += "ㅎ"
		} else {
			result += s
		}
	}

	return result
}
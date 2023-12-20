package main

import "testing"

// getChosung 함수가 잘 작동하는지 테스트하는 함수
func Test_getChosung(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{{
		in:   "생크림",
		want: "ㅅㅋㄹ",
	}, {
		in:   "생ㅋㄹ",
		want: "ㅅㅋㄹ",
	}, {
		in:   "생크림 라떼",
		want: "ㅅㅋㄹ ㄹㄸ",
	},
	}
	for _, c := range cases {
		b := getChosung(c.in)
		if c.want != b {
			t.Fatalf("Test_getChosung(): 입력 값: %v, 원하는 값: %v, 얻은 값: %v\n", c.in, c.want, b)
		}
	}
}
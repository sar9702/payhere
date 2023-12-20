# payhere

### 폴더구조

```
payhere
 ┣ assets
 ┃ ┣ js         // Javascript 파일
 ┃ ┃ ┣ jquery-3.7.1.min.js
 ┃ ┃ ┗ payhere.js
 ┃ ┗ templates  // HTML 파일
 ┃ ┃ ┣ detail.html
 ┃ ┃ ┣ edit-success.html
 ┃ ┃ ┣ edit.html
 ┃ ┃ ┣ footer.html
 ┃ ┃ ┣ header.html
 ┃ ┃ ┣ init.html
 ┃ ┃ ┣ navbar.html
 ┃ ┃ ┣ page.html
 ┃ ┃ ┣ register-success.html
 ┃ ┃ ┣ register.html
 ┃ ┃ ┣ signin.html
 ┃ ┃ ┣ signup-success.html
 ┃ ┃ ┗ signup.html
 ┣ README.md
 ┣ common.go
 ┣ common_test.go
 ┣ crypto.go
 ┣ db.sql
 ┣ dbapi.go
 ┣ dbapi_user.go
 ┣ go.mod
 ┣ go.sum
 ┣ handle_item.go
 ┣ handle_user.go
 ┣ main.go
 ┣ middleware.go
 ┣ payhere
 ┣ restapi_item.go
 ┣ struct.go
 ┗ testRun.sh
```

- handle\_\*.go : 핸들러 설정
- restapi_item.go : 아이템에 관한 Rest API
- dbapi\*.go : DB 관련 스크립트
- db.sql : DB DDL 파일
- testRun.sh : 서비스 실행 스크립트

<br>

### Rest API

|      URI      | Method |          Description          |     Parameter      |
| :-----------: | :----: | :---------------------------: | :----------------: |
|  /api/items   |  GET   |  모든 아이템 리스트 가져오기  | searchword, cursor |
|   /api/item   |  POST  |          아이템 등록          |        dept        |
| /api/item/:id |  GET   | ID에 해당하는 아이템 가져오기 |
| /api/item/:id |  PUT   | ID에 해당하는 아이템 업데이트 |
| /api/item/:id | DELETE |   ID에 해당하는 아이템 삭제   |

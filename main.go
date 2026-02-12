package main

import (
	"database/sql"
	"net/http"

	"github.com/doytowin/goooqo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./test.db")

	UserDataAccess := goooqo.NewDataAccess[UserEntity](db)
	goooqo.BuildRestService[UserEntity, UserQuery]("/user/", UserDataAccess)

	http.ListenAndServe(":9090", nil)
}

type UserEntity struct {
	goooqo.IntId
	Name  *string `json:"name,omitempty"`
	Score *int    `json:"score,omitempty"`
	Memo  *string `json:"memo,omitempty"`
}

type UserQuery struct {
	goooqo.PageQuery
}

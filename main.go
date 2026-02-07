package main

import (
	"database/sql"
	"net/http"

	"github.com/doytowin/goooqo/core"
	"github.com/doytowin/goooqo/rdb"
	"github.com/doytowin/goooqo/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./test.db")

	UserDataAccess := rdb.NewDataAccess[UserEntity](db)
	web.BuildRestService[UserEntity, UserQuery]("/user/", UserDataAccess)

	http.ListenAndServe(":9090", nil)
}

type UserEntity struct {
	core.Int64Id
	Name  *string `json:"name,omitempty"`
	Score *int    `json:"score,omitempty"`
	Memo  *string `json:"memo,omitempty"`
}

type UserQuery struct {
	core.PageQuery
}

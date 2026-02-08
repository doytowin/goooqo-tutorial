---
theme: custom
class: invert
size: 16:9
backgroundColor: #40b6e0
color: #fff
footer: 'https://github.com/doytowin/goooqo © 2026 DoytoWin, Inc'
---

# A Quick Start for GoooQo

### More than just an ORM for Golang

<br>

<div style="width: 85%; text-align: right;">—— f0rb</div>

---

## Table: t_user

<table><thead><tr><th>id</th><th>name</th><th data-type="number">score</th><th>memo</th></tr></thead><tbody><tr><td>1</td><td>Alley</td><td>80</td><td>Good</td></tr><tr><td>2</td><td>Dave</td><td>75</td><td>Well</td></tr><tr><td>3</td><td>Bob</td><td>60</td><td></td></tr><tr><td>4</td><td>Tim</td><td>92</td><td>Great</td></tr><tr><td>5</td><td>Emy</td><td>100</td><td>Great</td></tr><tr><td colspan="4">...</td></tr><tr><td colspan="4"></td></tr></tbody></table>

---

### UserEntity, UserQuery and Initialization

```go
type UserEntity struct {
    core.Int64Id
    Name  *string `json:"name,omitempty"`
    Score *int    `json:"score,omitempty"`
    Memo  *string `json:"memo,omitempty"`
}

type UserQuery struct {
    core.PageQuery
    //...
}
```

```go
UserDataAccess := rdb.NewDataAccess[UserEntity](db)
web.BuildRestService[UserEntity, UserQuery]("/user/", UserDataAccess)
```

---

## REST API

```shell
curl http://localhost:9090/user/
curl http://localhost:9090/user/?size=5
curl http://localhost:9090/user/?size=5&page=2
curl http://localhost:9090/user/?size=5&sort=memo&sort=score,desc
curl http://localhost:9090/user/1
curl -X DELETE http://localhost:9090/user/1
curl -X POST http://localhost:9090/user/ -d '[{"name":"John","score":85,"memo":"Good"}]'
curl -X PUT http://localhost:9090/user/21 -d '{"name":"John","score":95,"memo":"Great"}'
curl -X PATCH http://localhost:9090/user/21 -d '{"score":92}'
```

---

## Dynamic Queries

```shell
## Query all data where the memo is "Good".
curl http://localhost:9090/user/?memo=Good
## Query data with score ranging from 80 to 90.
curl http://localhost:9090/user/?scoreGe=80&scoreLt=90
## Query data where the memo is null.
curl http://localhost:9090/user/?memoNull=true
## Query data where the memo is not null.
curl http://localhost:9090/user/?memoNull=false
## Update the data where memo is "Good" to "OK".
curl -X PATCH http://localhost:9090/user/?memo=Good -d '{"memo":"OK"}'
## Delete the data where memo is "OK".
curl -X DELETE http://localhost:9090/user/?memo=OK
```

---

##

![width:1024px](https://goooqo.docs.doyto.win/~gitbook/image?url=https%3A%2F%2F2953169758-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FcaxovyAjMzGHxhZgzo8F%252Fuploads%252FYdPgvanlGOUBWvdn4L3R%252Fimage.png%3Falt%3Dmedia%26token%3Db996061d-d974-48fc-8202-45998b7545ef&width=768&dpr=2&quality=100&sign=efe50fab&sv=2)

---

## Next Video

- AND/OR conditions
- Subquery conditions


GoooQo Tutorial
---

- [Quick Start](assets/goooqo.md)

Run the `main` method to start the web server and execute the following commands to visit the REST APIs:

```shell
## Query the data list.
curl http://localhost:9090/user/
## Query the first 5 records.
curl http://localhost:9090/user/?size=5
## Query the second page.
curl http://localhost:9090/user/?size=5&page=2
## Sorting the first 5 records by memo, score.
curl http://localhost:9090/user/?size=5&sort=memo&sort=score,desc
## Query a record by id.
curl http://localhost:9090/user/1
## Delete a record by id.
curl -X DELETE http://localhost:9090/user/1
## Insert new records.
curl -X POST http://localhost:9090/user/ -d '[{"name":"John","score":85,"memo":"Good"}]'
## Update a record by id.
curl -X PUT http://localhost:9090/user/21 -d '{"name":"John","score":95,"memo":"Great"}'
## Update non-null fields by id.
curl -X PATCH http://localhost:9090/user/21 -d '{"score":92}'
```

Add the following fields to `UserQuery` to test dynamic queries:

```go
type UserQuery struct {
	goooqo.PageQuery
	Memo     *string
	ScoreLt  *int
	ScoreGe  *int
	MemoNull *bool
}
```

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

## Video on YouTube

[![Quick Start for GoooQo](https://img.youtube.com/vi/Tv2cKbMCvBQ/0.jpg)](https://youtu.be/Tv2cKbMCvBQ?si=EH6RiBf18ovLt7tx)


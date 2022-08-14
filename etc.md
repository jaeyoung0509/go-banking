- postgres는 기본적으로 로컬 접속을 신뢰하기 때문에 password를 묻지않음
- `golang-migrate`를 이용한 db migration
- - postgres container doesn’t enable SSL by default
```shell
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 
```
- `docker exec -it  container_name /bin/sh` (표준 리눅스 쉘 cli commands 사용가능)

- makefile 이란
- - Linux상에서 반복 적으로 발생하는 컴파일을 쉽게하기위해서 사용 


#### golang query mapper and orm
- sqlx 
- sqlx
- - failure won't occur until runtime
- sqlc
- - very fast and easy to use
- - catch sql query errors before generating codes 

#### sqlc
- generate (최소 한개의 쿼리가 있어야 됨)
- -  models.go 
- -  db.go 
- - 해당query.sql.go 
- 장점이자 단점: 생성된 파일의 내용을 수정해서는 안 됨
- `exec`: 반환 x 

#### unittesting 
- `testify` 사용하여 assertion 


#### why do we need db transaction ? (acid)
- to provide a reliable and consistent uow, even in case of system failure 
- to provide isolation between programs that access the database concurrently
- implement method 
- - define **store struct**


##### transaction lock & how to handle deadlcok
```sql
SELECT * FROM accounts WHERE id = 1;
SELECT * FROM accounts WHERE id = 1 FOR UPDATE; //query with lock

```

#### debug a deadlcok
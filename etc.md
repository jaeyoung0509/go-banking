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

#### avoid deadlock
- the best defense against deadlocks is to avoid them by making sure that
 our application always acquire locks in a consistent order.

#### understand isolation levels & read phenomena in db  (현상)
- there are four isolation level 
- `dirty read`phenomenon
  -  It happens when a transaction reads data written by other concurrent transaction that has not been committed yet
- `non-repeatable read`phenomenon
  - when transaction reads the same record twice and see different values
  because the row has been modified by other transaction that was committed after ther first read 
- `phantom read`phenomenon
  - it is similar phenomenon, but affects queries that search for multiple rows instead of one 
  
- `serializable anomaly`
  - when the result of a group of concurrent commited transactions could not be achived if we try to run them sequentially in any order without overlapping each other 
  
|                       | read uncommited | read commited | repeatable read | serializable |
|-----------------------|-----------------|---------------|-----------------|--------------|
| dirty read            | O               | X             | X               | X            |
| non-repetable read    | O               | O             | X               | X            |
| phantom read          | O               | O             | X               | X            |
| serialization anomaly | O               | O             | O               | X            |


#### github action 
- job 
  - is a set of steps execute on the same runner 
  - normal jobs run in parallel
- steps
  - is an individual task
  - run serially within a job 
  - contain 1 + actions 
- actions 
  - is a standalone command
  - run serially within a steps
  - can be reused
-  workflow
   -  event trigger
   -  scheduled trigger
   -  manually trigger 
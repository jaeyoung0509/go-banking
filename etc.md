- postgres는 기본적으로 로컬 접속을 신뢰하기 때문에 password를 묻지않음
- `golang-migrate`를 이용한 db migration
- - postgres container doesn’t enable SSL by default
```shell
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 
```
- `docker exec -it  container_name /bin/sh` (표준 리눅스 쉘 cli commands 사용가능)

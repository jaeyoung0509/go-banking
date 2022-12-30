## go banking 

### for learning
- [x] transaction management (acid)
  - with go channel
- [x] acheive test coveage 100 %
  - mocking (go mocking)
  - testing (testify)
- [ ] aws service 
- [ ] auth
- [ ] kubernetes  basic
- [ ] ssl tls 
- [ ] grpc basic 




---
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
  - 좀 더  단순하게 테스트 가능

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
 - select for update 
 - select no key for update 

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

#### manage env  
- using `viper`


#### mocking db 
**TODO can not generate mockgen command in makefile**
- it helps you to write independent tests more easily  
- tests will run much faster since they don't have to spend time talking to the db 
- it allows us to write test that acheive 100% coverage 
`mockgen --build_flags=--mod=mod  github.com/jaeyoung0509/go-banking/db/sqlc Store`

### secure password 
- use `bcrpyt`
  - bcrpyt requires a cost parameter which will decide the number of key expansion rounds or iterations of the algorithm
  - bcrpt also generates a random salt to be used in those iterations, which will help protect against the `rainbow table attack`
  - because of this `random salt` the algorithm will give you a completely different output hash baluse even if the same input password is provided
  - The cost and salt will also be added to the hash to produce the final hash string
  - in this hash string, there are 4 components
    - the first part is the `hash algorithm identifier` 
    - the second part is cost 
    - the third is the salt of length 16 byte
    - the last part is 24 bytes hash value


#### paseto
- what is`pasto` 
  - stronger algorithms
  - local  - symmetric key 
- what's the problem of jwt 
  - weak algorithms
  given developers too many algorithms to choose
  some algorithms are known to be vulnerable
  - trivial forgery
  - the developers should check the algorithm header before verify toekn signature
  server must check if the algorithm header matches the one it is using to sign tokens


#### [db] 트랜잭션 격리 수준
- 트랜잭션 수준 읽기 일관성 (동시성 제어)를 위
```
- READ UNCOMMITTED : 다른 트랜잭션에 커밋되지 않은 내용도 참조 가능
- READ COMMITTED: 다른 트랜잭션에 커밋된 내용만 참조
- REPEATABLE READ: 트랜잭션 진입하기 이전에 커밋된 내용만 참조
- SERIALIZABLE: 트랜잭션에 진입하면 락을 걸어 다른 트랜잭션이 접근하지 못하게or
```
- 아래로 내려갈수록 트랜잭션간 고립 정도가 높아지며, 성능이 떨어짐
- 일반적인 온라인 서비스에서는 read-commited or repeatable read를 사용


### github action
![](2022-12-27-14-50-13.png)
#### runner
- is a server to run the jobs
- run 1 job at a time
- github hosted or self hosted
- report progress, logs & result to github
#### job
- is a set of steps execute on the same runner 
- normal jobs in parallel
- dependent jobs run serially
#### step
- is an individual task
- run serially within a job
- contain 1  + actions 
#### action
- is a standalone command
- run serially within a step
- can be reused


#### why mock database
- 독립적인 테스트
  - 테스트 데이터 셋 충돌을 피할 수 있음
- 빠른 테스트
  - 디비 연결하는 시간 감소 (?)

#### how to mock
- use fake db: memory
- use db stubs: gomock
  - generate and build stubs that returns hard-coded values


#### symmetric <-> asymmetric
- symmetric 
  - the same secret key 
  - for local use 
  - hs
    - hs256 = hmac + sha256

- asymmetric 
  - the private key is used to sing token
  - the public key is used to verify token
  - for public use: internal service token
  but external service needs to verify it
  - rs256: rsa pcks + sha256
    - public key cryptography standards
  - ps256: rsa pss + sha256
    - pss: probabilistic signature scheme
  - es256:
    - ecdsa: elliptic curve digial signature algorithm

- the problem of jwt
  - weak algorithm
    - rsa pkcs: padding oracle attack
    - ecdsa: invalid curve attack
  - include the signing algorithm in the token header 
    - algon = none
    - it fixed many libraries
  -  set 'alg' header to 'hs256' while the server normally verifies token with a rsa public key 
     -  even though the token is rsa
     but hacker set algo as 'hs256'(symmetric) then server will use symmetric method instead of rsa public
     - so server should check token's algorithm
#### paseto
  - platform-agnostic-security-tokens
  - stronger algorithms
  - paseto have two algorithms
    - symmetric
    - asymmetric
  - jwt use base64 <-> paseto aead
    - ssl/ tls
    - aead use crt with hmac sha384
   ![](2022-12-30-22-56-17.png)
   ![](2022-12-30-22-58-49.png)
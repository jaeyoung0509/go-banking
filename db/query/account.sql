-- name: CreateAccount :one
insert into accounts (
    owner,
    balance,
    currency
) values (
    $1 , $2, $3
) returning *;


-- name: GetAccount :one
select * from accounts 
where id = $1 limit 1;

-- name: ListAccounts :many
select * from accounts 
order by id
limit $1 
offset $2;

-- name: UpdateAccount :one
update accounts
set balance =$2 
where id = $1
returning *;


-- name: DeleteAccounts :exec
delete from accounts 
where id =$1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;



-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;
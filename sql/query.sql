-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :exec
INSERT INTO users (
  username, password
) VALUES (
  ?, ?
);

-- name: GetUser :one
select * from users
where username == ?;
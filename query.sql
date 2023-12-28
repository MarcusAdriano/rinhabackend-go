-- name: GetPessoa :one
SELECT *
FROM pessoas
WHERE id = $1 LIMIT 1;

-- name: CreatePessoa :one
INSERT INTO pessoas (id, nome, apelido, stack)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CountPessoas :one
SELECT COUNT(*)
FROM pessoas;

-- name: SearchPessoas :many
SELECT *
FROM pessoas
WHERE nome ILIKE $1 OR apelido ILIKE $1 OR stack ILIKE $1;

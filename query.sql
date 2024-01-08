-- name: GetPessoa :one
SELECT id,nome,apelido,stack,nascimento
FROM pessoas
WHERE id = $1 LIMIT 1;

-- name: CreatePessoa :exec
INSERT INTO pessoas (id, nome, apelido, stack, nascimento)
VALUES ($1, $2, $3, $4, $5);

-- name: CountPessoas :one
SELECT COUNT(*)
FROM pessoas;

-- name: SearchPessoas :many
SELECT id,nome,apelido,stack,nascimento
FROM pessoas
WHERE busca @@ plainto_tsquery($1)
ORDER BY nascimento DESC
LIMIT 10;

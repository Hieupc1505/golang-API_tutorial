-- name: CreateTodoTask :one
INSERT INTO todo_list (
    title,
    image,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTodoTask :one
SELECT * FROM todo_list
WHERE id = $1 LIMIT 1;

-- name: GetTodoTaskList :many
SELECT * FROM todo_list
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTodoTask :one
UPDATE todo_list
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTodoTask :exec
DELETE FROM todo_list
WHERE id = $1;
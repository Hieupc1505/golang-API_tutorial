// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: todo_list.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTodoTask = `-- name: CreateTodoTask :one
INSERT INTO todo_list (
    title,
    image,
    description
) VALUES (
    $1, $2, $3
) RETURNING id, title, image, description, status, created_at, updated_at, "user"
`

type CreateTodoTaskParams struct {
	Title       string      `json:"title"`
	Image       pgtype.Text `json:"image"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) CreateTodoTask(ctx context.Context, arg CreateTodoTaskParams) (TodoList, error) {
	row := q.db.QueryRow(ctx, createTodoTask, arg.Title, arg.Image, arg.Description)
	var i TodoList
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Image,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.User,
	)
	return i, err
}

const deleteTodoTask = `-- name: DeleteTodoTask :exec
DELETE FROM todo_list
WHERE id = $1
`

func (q *Queries) DeleteTodoTask(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTodoTask, id)
	return err
}

const getTodoTask = `-- name: GetTodoTask :one
SELECT id, title, image, description, status, created_at, updated_at, "user" FROM todo_list
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodoTask(ctx context.Context, id int64) (TodoList, error) {
	row := q.db.QueryRow(ctx, getTodoTask, id)
	var i TodoList
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Image,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.User,
	)
	return i, err
}

const getTodoTaskList = `-- name: GetTodoTaskList :many
SELECT id, title, image, description, status, created_at, updated_at, "user" FROM todo_list
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetTodoTaskListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetTodoTaskList(ctx context.Context, arg GetTodoTaskListParams) ([]TodoList, error) {
	rows, err := q.db.Query(ctx, getTodoTaskList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TodoList{}
	for rows.Next() {
		var i TodoList
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Image,
			&i.Description,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.User,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodoTask = `-- name: UpdateTodoTask :one
UPDATE todo_list
SET status = $2
WHERE id = $1
RETURNING id, title, image, description, status, created_at, updated_at, "user"
`

type UpdateTodoTaskParams struct {
	ID     int64      `json:"id"`
	Status TaskStatus `json:"status"`
}

func (q *Queries) UpdateTodoTask(ctx context.Context, arg UpdateTodoTaskParams) (TodoList, error) {
	row := q.db.QueryRow(ctx, updateTodoTask, arg.ID, arg.Status)
	var i TodoList
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Image,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.User,
	)
	return i, err
}

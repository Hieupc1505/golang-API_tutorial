package db

import (
	"context"
	"testing"
	"time"

	"github.com/hieupc09/simple_api/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomTodoTask(t *testing.T) TodoList {
	arg := CreateTodoTaskParams{
		Title:       utils.RandomString(6),
		Description: pgtype.Text{String: utils.RandomDescription(), Valid: true},
	}

	todo, err := testQueries.CreateTodoTask(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, arg.Description, todo.Description)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreatedAt)
	require.NotZero(t, todo.Status)

	return todo

}

func TestCreateTodo(t *testing.T) {
	createRandomTodoTask(t)
}

func TestGetTodoTask(t *testing.T) {
	task1 := createRandomTodoTask(t)
	task2, err := testQueries.GetTodoTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Title, task2.Title)
	require.Equal(t, task1.Description, task2.Description)
	require.Equal(t, task1.Status, task2.Status)
	require.WithinDuration(t, task1.CreatedAt.Time, task2.CreatedAt.Time, time.Second)
}

func TestUpdateTodoTask(t *testing.T) {
	task1 := createRandomTodoTask(t)
	// Sử dụng RandomEnum để chọn ngẫu nhiên một TaskStatus
	randomStatus := utils.RandomEnum([]TaskStatus{TaskStatusDelete, TaskStatusDoing, TaskStatusDone}).(TaskStatus)
	arg := UpdateTodoTaskParams{
		ID:     task1.ID,
		Status: randomStatus,
	}
	task2, err := testQueries.UpdateTodoTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Title, task2.Title)
	require.Equal(t, task1.Description, task2.Description)
	require.Equal(t, arg.Status, task2.Status)
	require.WithinDuration(t, task1.CreatedAt.Time, task2.CreatedAt.Time, time.Second)
}

func TestGetTodoTaskList(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTodoTask(t)
	}
	arg := GetTodoTaskListParams{
		Limit:  5,
		Offset: 5,
	}
	todos, err := testQueries.GetTodoTaskList(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, todos, 5)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
	}
}

func TestDeleteTodoTak(t *testing.T) {
	task1 := createRandomTodoTask(t)
	err := testQueries.DeleteTodoTask(context.Background(), task1.ID)
	require.NoError(t, err)

	task2, err := testQueries.GetTodoTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error()) //pgx.ErrNoRows.Error()
	require.Empty(t, task2)
}

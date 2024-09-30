package db

import (
	"context"
	"testing"

	"github.com/hieupc09/simple_api/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hassPass, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	args := CreateUserParams{
		Username: utils.RandomOwner(),
		Password: hassPass,
		Email:    utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Password, user.Password)
	require.Equal(t, args.Email, user.Email)

	require.NotZero(t, user.ChangePasswordAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ChangePasswordAt.Time, user2.ChangePasswordAt.Time)
	require.Equal(t, user1.CreatedAt.Time, user2.CreatedAt.Time)
}

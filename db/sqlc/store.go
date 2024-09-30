package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(coon *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  New(coon),
		connPool: coon,
	}
}

package pg

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}
func New(once *sync.Once, ctx context.Context) (*Postgres, error) {
	pg:= &Postgres{}
	var err error
	// TODO: make them configurable
	connStr := "postgres://postgres:postgres@localhost:5432/dev?"
	once.Do(func() {
		pg.pool, err = pgxpool.New(ctx, connStr)
	})
	if err != nil {
		return pg, err
	}
	return pg, nil
}


func (pg *Postgres) GetConn(ctx context.Context) (*pgxpool.Conn,error) {
	return pg.pool.Acquire(ctx)
}


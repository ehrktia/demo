package repo

import (
	"context"
	"errors"

	"github.com/ehrktia/demo/entity"
	"github.com/ehrktia/demo/store/pg"
	"github.com/jackc/pgx/v5"
)

type UserStore struct {
	connPool *pg.Postgres
}

func New(connPool *pg.Postgres) *UserStore {
	return &UserStore{
		connPool: connPool,
	}
}

var ErrDuplicateData = errors.New("duplicate user found")
var NoDataFound = errors.New("no data found")

func (us *UserStore) GetUserById(ctx context.Context, id int) (entity.User, error) {
	conn, err := us.connPool.GetConn(ctx)
	if err != nil {
		return entity.User{}, err
	}
	defer func() {
		conn.Conn().Close(ctx)
		conn.Release()
	}()
	query := `select * from users where user_id=$1`
	rows, err := conn.Query(ctx, query, id)
	if err != nil {
		return entity.User{}, err
	}
	users, err := pgx.CollectRows(rows, pgx.RowTo[entity.User])
	if err != nil {
		return entity.User{}, err
	}
	if len(users) > 1 {
		return entity.User{}, ErrDuplicateData
	}
	if len(users) < 1 {
		return entity.User{}, NoDataFound
	}
	return users[0], nil
}

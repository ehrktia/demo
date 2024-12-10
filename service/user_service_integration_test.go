//go:build integration

package service

import (
	"sync"
	"testing"

	"github.com/ehrktia/demo/repo"
	"github.com/ehrktia/demo/store/pg"
)

func Test_get_user(t *testing.T) {
	once := &sync.Once{}
	conn, err := pg.New(once, ctx)
	if err != nil {
		t.Fatal(err)
	}
	ur := repo.New(conn)
	us := NewUserService(ur)
	id := 1
	if _, err := us.GetUserById(ctx, id); err == nil {
		t.Fatal(err)
	}
}

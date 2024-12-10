package service

import (
	"context"
	"testing"

	"github.com/ehrktia/demo/entity"
	"github.com/ehrktia/demo/mocks"
	"github.com/ehrktia/demo/repo"
)

var ctx = context.Background()

func Test_user_get(t *testing.T) {
	t.Run("sucessful get user", func(t *testing.T) {
		id := 1
		mockUserRepo := mocks.NewUserRepository(t)
		u := entity.User{
			Id: id, Name: t.Name(),
		}
		mockUserRepo.EXPECT().GetUserById(ctx, id).Return(u, nil)
		us := NewUserService(mockUserRepo)
		got, err := us.GetUserById(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		if got.Id != id {
			t.Fatalf("expected:%d,got:%d", id, got.Id)
		}
	})
	t.Run("error fetching user from store", func(t *testing.T) {
		mockUserRepo := mocks.NewUserRepository(t)
		id := 1
		u := entity.User{
			Id: 1, Name: t.Name(),
		}
		mockUserRepo.EXPECT().GetUserById(ctx, id).Return(u, repo.NoDataFound)
		us := NewUserService(mockUserRepo)
		if _, err := us.GetUserById(ctx, id); err == nil {
			t.Fatalf("expected:%s,got:%s", repo.NoDataFound.Error(), err.Error())

		}

	})

}

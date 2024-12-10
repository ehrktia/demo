package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/ehrktia/demo/repo"
	"github.com/ehrktia/demo/service"
	"github.com/ehrktia/demo/store/pg"
	"github.com/ehrktia/demo/web"
)

func main() {
	once := &sync.Once{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	connPool, err := pg.New(once, ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create conn pool:%v", err)
	}
	userStore := repo.New(connPool)
	us := service.NewUserService(userStore)
	port := "8080"
	webServer := web.NewWebServer(us, port)
	webServer.RegisterRoutes()
	log := slog.Default()
	log.Info("webserver started successfully", slog.Any("port", port))
	if err := webServer.HTTPServer().ListenAndServe(); err != nil {
		panic(err)
	}

}

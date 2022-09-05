//go:build wireinject
// +build wireinject

package main

import (
	"demo/internal/app"
	"demo/internal/pkg/infra"
	"demo/internal/pkg/repository"
	"demo/internal/pkg/service"
	"demo/internal/server/http"
	"github.com/google/wire"
)

func InitializeServer() *http.Server {
	panic(wire.Build(
		http.NewHttpServer,
		app.NewApp,
		service.NewDeviceUseCase,
		repository.NewDevicePostgresRepo, repository.NewDeviceRedisRepo,
		infra.NewDBConn, infra.NewRedisConn,
	))
}

//go:build wireinject
// +build wireinject

package di

import (
	http "Clean-Grpc/pkg/api"
	"Clean-Grpc/pkg/api/handler"
	"Clean-Grpc/pkg/config"
	"Clean-Grpc/pkg/db"
	"Clean-Grpc/pkg/repository"
	"Clean-Grpc/pkg/usecase"

	"github.com/google/wire"
)


func  InitializeAPI(cfg config.Config) (*http.ServerHTTP, error){
	wire.Build(
		db.ConnectDatabase,

		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
	    http.NewServerHTTP)

    return&http.ServerHTTP{},nil
}
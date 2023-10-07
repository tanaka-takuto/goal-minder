//go:build wireinject
// +build wireinject

package di

import (
	"goal-minder/adapter/repository"
	"goal-minder/domain/usecase"

	"github.com/google/wire"
)

func CreateAccountUsecase() usecase.CreateAccountUsecase {
	wire.Build(
		repository.RepositorySet,
		wire.Struct(new(usecase.CreateAccountUsecase), "*"),
	)
	return usecase.CreateAccountUsecase{}
}

func LoginUsecase() usecase.LoginUsecase {
	wire.Build(
		repository.RepositorySet,
		wire.Struct(new(usecase.LoginUsecase), "*"),
	)
	return usecase.LoginUsecase{}
}

func MeUsecase() usecase.MeUsecase {
	wire.Build(
		repository.RepositorySet,
		wire.Struct(new(usecase.MeUsecase), "*"),
	)
	return usecase.MeUsecase{}
}

func SetGoalUsecase() usecase.SetGoalUsecase {
	wire.Build(
		repository.RepositorySet,
		wire.Struct(new(usecase.SetGoalUsecase), "*"),
	)
	return usecase.SetGoalUsecase{}
}

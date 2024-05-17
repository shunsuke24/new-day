package usecase

import (
	"github.com/new-day/domain/model"
	"github.com/new-day/env"
	"github.com/new-day/infra"
	"github.com/new-day/infra/redis"
	"github.com/new-day/repository"
)

type Usecase interface {
	Recommend() (string, error)
	CreateUser(user *model.User) (string, error)
}

type usecase struct {
	openai   *repository.OpenAI
	conf     *env.Config
	redis    *infra.Redis
	userRepo redis.UserRepository
}

type Params struct {
	OpenAI   *repository.OpenAI
	Conf     *env.Config
	Redis    *infra.Redis
	UserRepo redis.UserRepository
}

func NewUsecase(p *Params) Usecase {
	return &usecase{
		openai:   p.OpenAI,
		conf:     p.Conf,
		redis:    p.Redis,
		userRepo: p.UserRepo,
	}
}

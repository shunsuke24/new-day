package usecase

import (
	"github.com/new-day/env"
	"github.com/new-day/repository"
)

type Usecase interface {
	Recommend() (string, error)
}

type usecase struct {
	openai *repository.OpenAI
	conf   *env.Config
}

type Params struct {
	OpenAI *repository.OpenAI
	Conf   *env.Config
}

func NewUsecase(p *Params) Usecase {
	return &usecase{
		openai: p.OpenAI,
		conf:   p.Conf,
	}
}

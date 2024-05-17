package usecase

import (
	"github.com/new-day/domain/model"
)

func (u *usecase) CreateUser(user *model.User) (string, error) {
	err := u.userRepo.Save(user)
	if err != nil {
		return "fail", err
	}
	return "success", nil
}

package redis

import (
	"encoding/json"
	"fmt"

	"github.com/new-day/domain/model"
	"github.com/new-day/infra"
)

type UserRepository struct {
	redis *infra.Redis
}

func (r *UserRepository) Save(user *model.User) error {
	key := fmt.Sprintf("user:%d", user.ID)
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.redis.Set(key, value)
}

// func (r *UserRepository) FindByID(id int) (*model.User, error) {
// 	key := fmt.Sprintf("user:%d", id)
// 	value, err := r.redis.Get(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if value == nil {
// 		return nil, nil
// 	}
// 	var user model.User
// 	err = json.Unmarshal(value, &user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

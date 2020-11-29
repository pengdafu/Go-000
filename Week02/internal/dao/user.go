package dao

import (
	"github.com/pkg/errors"
	"week02/internal/model"
)

type UserDao struct{}

type User struct {
	ID       int
	TrueName string
}

func (UserDao) FindUserById(id int) (User, error) {
	u := User{}
	err := model.DB().Table("t_user").Where("id = ?", id).Find(&u).Error
	return u, errors.Wrapf(err, "find user error %v", err)
}

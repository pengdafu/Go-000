package dao

import "github.com/pkg/errors"

type User struct {
	ID       int
	TrueName string
}

func FindUserById(id int) (User, error) {
	u := User{}
	err := db.Table("t_user").Where("id = ?", id).Find(&u).Error
	return u, errors.Wrapf(err, "find user error %v", err)
}

package services

import (
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/model/users"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RESTErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

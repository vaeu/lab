package services

import (
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/model/users"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

func GetUser(uID uint64) (*users.User, *errors.RESTErr) {
	result := &users.User{ID: uID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(u users.User) (*users.User, *errors.RESTErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

func UpdateUser(u users.User) (*users.User, *errors.RESTErr) {
	current, err := GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	current.FirstName = u.FirstName
	current.LastName = u.LastName
	current.Email = u.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

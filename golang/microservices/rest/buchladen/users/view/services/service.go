package services

import "github.com/vaeu/lab/golang/microservices/rest/buchladen/users/model/users"

func CreateUser(u users.User) (*users.User, error) {
	return &u, nil
}

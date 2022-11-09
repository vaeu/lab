package users

import (
	"strings"

	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

type User struct {
	ID          uint64 `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (u *User) Validate() *errors.RESTErr {
	email := strings.TrimSpace(strings.ToLower(u.Email))
	if email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	return nil
}

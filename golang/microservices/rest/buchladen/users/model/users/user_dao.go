package users

import (
	"fmt"

	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

var usersDB = make(map[uint64]*User)

func (u *User) Get() *errors.RESTErr {
	res := usersDB[u.ID]
	if res == nil {
		errors.NewNotFound(fmt.Sprintf("user not found: %d", u.ID))
	}

	u.ID = res.ID
	u.FirstName = res.FirstName
	u.LastName = res.LastName
	u.Email = res.Email
	u.DateCreated = res.DateCreated

	return nil
}

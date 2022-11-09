package users

import (
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

var usersDB = make(map[uint64]*User)

func (u *User) Get() *errors.RESTErr {
	// get UID
	return nil
}

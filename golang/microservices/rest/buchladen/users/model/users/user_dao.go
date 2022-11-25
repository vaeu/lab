package users

import (
	"fmt"
	"strings"

	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/db/mysql/usersdb"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/dates"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var usersDB = make(map[uint64]*User)

func (u *User) Get() *errors.RESTErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

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

func (u *User) Save() *errors.RESTErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	defer stmt.Close()
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}

	u.DateCreated = dates.GetNowString()

	insertResult, err := stmt.Exec(
		u.FirstName, u.LastName, u.Email, u.DateCreated,
	)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequest(fmt.Sprintf("email address is already taken: %s", u.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("unable to save user: %s\n", err.Error()),
		)
	}

	uID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("unable to save user: %s\n", err.Error()),
		)
	}
	u.ID = uint64(uID)

	return nil
}

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
	errNoRows        = "no rows in result set"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

func (u *User) Get() *errors.RESTErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	defer stmt.Close()
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}

	res := stmt.QueryRow(u.ID)
	if err := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFound(
				fmt.Sprintf("user %d does not exist", u.ID),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("unable to get UID %d: %s", u.ID, err.Error()),
		)
	}
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
			fmt.Sprintf("unable to save user: %s", err.Error()),
		)
	}

	uID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("unable to save user: %s", err.Error()),
		)
	}
	u.ID = uint64(uID)

	return nil
}

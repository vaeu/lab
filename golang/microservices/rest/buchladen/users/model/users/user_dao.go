package users

import (
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/db/mysql/usersdb"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/dates"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/mysqlutils"
)

const (
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

func (u *User) Get() *errors.RESTErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	defer stmt.Close()
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}

	res := stmt.QueryRow(u.ID)
	if err := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		return mysqlutils.ParseError(err)
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
		return mysqlutils.ParseError(err)
	}

	uID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	u.ID = uint64(uID)

	return nil
}

func (u *User) Update() *errors.RESTErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		mysqlutils.ParseError(err)
	}
	return nil
}

func (u *User) Delete() *errors.RESTErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(u.ID); err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

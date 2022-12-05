package mysqlutils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
)

const errNoRows = "no rows in result set"

func ParseError(err error) *errors.RESTErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFound("no rows found matching given ID")
		}
		return errors.NewInternalServerError("unable to parse db response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequest("duplicate entry")
	}
	return errors.NewInternalServerError("unable to process request")
}

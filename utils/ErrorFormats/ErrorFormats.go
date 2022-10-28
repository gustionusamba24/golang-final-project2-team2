package ErrorFormats

import (
	"golang-final-project2-team2/utils/ErrorUtils"
	"strings"
)

func ParseError(err error) ErrorUtils.MessageErr {

	if strings.Contains(err.Error(), "no rows in result set") {
		return ErrorUtils.NewNotFoundError("no record found")
	}
	return ErrorUtils.NewInternalServerError("something went wrong")
}

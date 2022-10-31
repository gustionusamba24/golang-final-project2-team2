package error_formats

import (
	"golang-final-project2-team2/utils/error_utils"
	"strings"
)

func ParseError(err error) error_utils.MessageErr {

	if strings.Contains(err.Error(), "no rows in result set") {
		return error_utils.NewNotFoundError("no record found")
	}
	return error_utils.NewInternalServerError("something went wrong")
}

func NoAuthorization() error_utils.MessageErr {
	return error_utils.NewUnauthorizedRequest("This user does not have authorization to perform action")
}

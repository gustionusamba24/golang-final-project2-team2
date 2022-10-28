package UserDomain

import (
	"golang-final-project2-team2/utils/ErrorUtils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,len=6"`
	Age       int8      `json:"age" validate:"required,min=9"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreateResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int8   `json:"age"`
}

func (u *User) Validate() ErrorUtils.MessageErr {
	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		return ErrorUtils.NewBadRequest(err.Error())
	}
	return nil
}

func (u *User) GetUserIdParam(c *gin.Context) (int64, ErrorUtils.MessageErr) {
	idParam := c.Param("userId")
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		return int64(0), ErrorUtils.NewBadRequest("invalid user id params")
	}

	return int64(userId), nil
}

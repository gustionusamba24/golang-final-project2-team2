package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-final-project2-team2/domains/user_domain"
	"golang-final-project2-team2/utils/error_utils"
	"golang.org/x/crypto/bcrypt"

	"os"
	"strings"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func HashPass(pass string) (*string, error_utils.MessageErr) {
	salt := 8
	password := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	hashString := string(hash)
	if err != nil {
		return nil, error_utils.NewInternalServerError(err.Error())
	}
	return &hashString, nil
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func GenerateToken(user *user_domain.User) (*string, error_utils.MessageErr) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, error_utils.NewInternalServerError(err.Error())

	}
	return &signedToken, nil
}

func VerifyToken(c *gin.Context) (interface{}, error_utils.MessageErr) {
	err := errors.New("please login to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, error_utils.NewUnauthorizedRequest(err.Error())
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, error_utils.NewInternalServerError(err.Error())
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, error_utils.NewInternalServerError(err.Error())
	}
	return token.Claims.(jwt.MapClaims), nil

}

func ValidateRequest(request interface{}) error_utils.MessageErr {
	validate := validator.New()

	err := validate.Struct(request)

	if err != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	return nil
}

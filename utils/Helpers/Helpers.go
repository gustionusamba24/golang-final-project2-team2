package Helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang-final-project2-team2/domains/UserDomain"
	"golang-final-project2-team2/utils/ErrorUtils"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func HashPass(pass string) (*string, ErrorUtils.MessageErr) {
	salt := 8
	password := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	hashString := string(hash)
	if err != nil {
		return nil, ErrorUtils.NewInternalServerError(err.Error())
	}
	return &hashString, nil
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func GenerateToken(user *UserDomain.User) (*string, ErrorUtils.MessageErr) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, ErrorUtils.NewInternalServerError(err.Error())

	}
	return &signedToken, nil
}

func VerifyToken(stringToken string) (interface{}, ErrorUtils.MessageErr) {
	err := errors.New("please login to proceed")
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(jwtSecretKey), nil
	})
	//if err != nil {
	//	return nil, ErrorUtils.NewInternalServerError(err.Error())
	//}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, ErrorUtils.NewInternalServerError(err.Error())
	}
	return token.Claims.(jwt.MapClaims), nil

}

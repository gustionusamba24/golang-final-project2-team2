package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"golang-final-project2-team2/domains/user_domain"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var (
	pass        = "rahasia"
	correctHash = "$2a$08$odrXKBHB8/XBe2lvEiPfdumgFthHm21cZzQWKlue74Pzyrs1EDzDm"
	wrongHash   = "$2a$08$odrXKBHB8/XBe2lvEiPfdumgFthHm21cZzQWKlue74Pzyrs1EDzDs"
	user        = &user_domain.User{
		Id:       1,
		Username: "testing",
		Email:    "testing@gmail.com",
	}
)

func TestSuccessGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(jwtSecretKey))

	if err != nil {
		t.Errorf("Terjadi kesalahan : %s", err.Error())
		return

	}
	t.Logf("Berhasil generate token : %s", signedToken)
}

func TestFailedGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString("")

	if err != nil {
		t.Errorf("Terjadi kesalahan : %s", err.Error())
		return

	}
	t.Logf("Berhasil generate token : %s", signedToken)
}

func TestSuccessComparePass(t *testing.T) {
	hash, pass := []byte(correctHash), []byte(pass)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		t.Errorf("Kedua password tidak cocok : %s", err.Error())
		return
	}
	t.Logf("Kedua password Cocok")
}

func TestFailedComparePass(t *testing.T) {
	hash, pass := []byte(wrongHash), []byte(pass)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		t.Errorf("Kedua password tidak cocok : %s", err.Error())
		return
	}
	t.Logf("Kedua password Cocok")
}

func TestSuccessHashPass(t *testing.T) {
	salt := 8
	password := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	hashString := string(hash)
	if err != nil {
		t.Errorf("Terjadi Kesalahan : %s", err.Error())
		return
	}
	t.Logf("Hasil Hash Password : %s", hashString)
}

func TestFailedHashPass(t *testing.T) {
	salt := 100
	password := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	hashString := string(hash)
	if err != nil {
		t.Errorf("Terjadi Kesalahan : %s", err.Error())
		return
	}
	t.Logf("Hasil Hash Password : %s", hashString)
}

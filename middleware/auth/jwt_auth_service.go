package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-pzn-clone/helper"
)

var SECRET_KEY = []byte("iamironmanforyou")

type JWTAuth interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTAuthImpl struct {
}

func NewJWTAuth() *JWTAuthImpl {
	return &JWTAuthImpl{}
}

func (j *JWTAuthImpl) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(SECRET_KEY)
	helper.PanicIfError(err)

	return signedString, nil
}

func (j *JWTAuthImpl) ValidateToken(token string) (*jwt.Token, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return SECRET_KEY, nil
	})

	if err != nil {
		return nil, errors.New("Kamu jahat")
	}

	return parse, nil
}

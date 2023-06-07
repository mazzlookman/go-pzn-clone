package test

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := jwtAuth.GenerateToken(1)
	if err != nil {
		t.Error(err)
	}

	splitToken := strings.Split(token, ".")
	assert.Equal(t, 3, len(splitToken))

	t.Log(token)

}

func TestValidateToken(t *testing.T) {
	token, err := jwtAuth.GenerateToken(2)
	if err != nil {
		t.Error(err)
	}

	t.Log(token)

	validateToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		t.Error(err)
	}

	claims := validateToken.Claims.(jwt.MapClaims)
	assert.Equal(t, 2, int(claims["user_id"].(float64)))

}

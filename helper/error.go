package helper

import (
	"errors"
	"go-pzn-clone/model/domain"
)

func ErrRepository(err error) (domain.User, error) {
	if err != nil {
		return domain.User{}, errors.New("Repository Error: " + err.Error())
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

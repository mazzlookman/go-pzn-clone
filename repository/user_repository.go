package repository

import "go-pzn-clone/model/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	FindByID(userID int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	DeleteByID(userID int) (bool, error)
}

package repository

import (
	"go-pzn-clone/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByID(userID int) (domain.User, error) {
	user := domain.User{}
	err := r.db.Where("id=?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Where("email=?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) DeleteByID(userID int) error {
	user := domain.User{}
	err := r.db.Delete(&user, "id=?", userID).Error
	if err != nil {
		return err
	}
	return nil
}

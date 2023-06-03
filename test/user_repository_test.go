package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"testing"
)

func TestSaveUser(t *testing.T) {
	user := domain.User{
		Name:     "Otong Lagi",
		Email:    "otonglagi@test.com",
		Password: "123",
		Role:     "user",
		Avatar:   "otong.jpg",
	}
	save, err := userRepo.Save(user)
	helper.PanicIfError(err)

	assert.Equal(t, 2, save.ID)
	assert.Equal(t, "Otong Lagi", save.Name)
	assert.Equal(t, "user", save.Role)
}

func TestFindByIDUser(t *testing.T) {
	findByID, err := userRepo.FindByID(1)
	helper.PanicIfError(err)

	//assert.Equal(t, "", findByID.Name)

	assert.Equal(t, 1, findByID.ID)
	assert.Equal(t, "Otong", findByID.Name)
	assert.Equal(t, "user", findByID.Role)
}

func TestFindByEmailUser(t *testing.T) {
	findByEmail, err := userRepo.FindByEmail("otong@test.com")
	helper.PanicIfError(err)

	//assert.Equal(t, 0, findByEmail.ID)

	assert.Equal(t, 1, findByEmail.ID)
	assert.Equal(t, "Otong", findByEmail.Name)
	assert.Equal(t, "user", findByEmail.Role)
}

func TestUpdate(t *testing.T) {
	findByID, err := userRepo.FindByID(1)
	helper.PanicIfError(err)

	findByID.Name = "Otong Keren"
	findByID.Avatar = "otongkeren.jpg"

	update, err := userRepo.Update(findByID)
	helper.PanicIfError(err)

	assert.Equal(t, "Otong Keren", update.Name)
	assert.Equal(t, 1, update.ID)
	assert.Equal(t, "otongkeren.jpg", update.Avatar)
}

func TestDeleteByIDUser(t *testing.T) {
	err := userRepo.DeleteByID(2)
	helper.PanicIfError(err)

	fmt.Println("User is deleted")
}

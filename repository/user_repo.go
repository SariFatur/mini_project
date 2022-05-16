package repository

import (
	"fmt"
	"myproject/model"

	"gorm.io/gorm"
)

type IUserRepo interface {
	GetAllUser() ([]model.UserCustom, error)
	GetUserByUsername(username string) (model.User, error)
	InsertUser(user model.User) error
	GetUserById(id int) (model.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db}
}

func (r UserRepo) GetAllUser() ([]model.UserCustom, error) {
	users := []model.UserCustom{}
	err := r.db.Find(&users).Error
	if err != nil {
		fmt.Println("Error While GetAllUser", err)
	}
	return users, err
}

func (repo UserRepo) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println("error while GetUserByUsernameAndPassword", err)
	}
	return user, err
}

func (repo UserRepo) InsertUser(user model.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		fmt.Println("error while InsertUser", err)
	}
	return err
}

func (r UserRepo) GetUserById(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println("error while GetUserById", err)
	}
	return user, err
}

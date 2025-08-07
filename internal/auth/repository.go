package auth

import (
	"errors"
	"exams/internal/modules/users"
	"exams/internal/utils"

	"gorm.io/gorm"
)

type Repository interface {
	Login(CheckUserLogin) (*users.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Login(req CheckUserLogin) (*users.User, error) {
	var user users.User

	if err := r.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

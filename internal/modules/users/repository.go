package users

import (
	"errors"
	"exams/internal/modules/teachers"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	GetUsersPaginated(offset int, limit int) ([]User, int64, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	CreateUserAndTeacher(user *User, teacher *teachers.Teacher) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&User{})
	return &repository{db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) CreateUserAndTeacher(user *User, teacher *teachers.Teacher) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		var count int64
		if err := tx.Model(&User{}).Where("Username = ?", user.Username).Count(&count).Error; err != nil {
			fmt.Print(err.Error())
			return err
		}

		if count > 0 {
			fmt.Print("username already taken")
			return errors.New("username already taken")
		}

		if err := tx.Create(user).Error; err != nil {
			return err
		}

		teacher.ID = user.ID
		if err := tx.Create(teacher).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUsersPaginated(offset, limit int) ([]User, int64, error) {
	var users []User
	var total int64

	if err := r.db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Preload("Teacher").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

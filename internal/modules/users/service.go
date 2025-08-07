package users

import (
	"exams/internal/modules/teachers"
	"exams/internal/utils"
)

type Service interface {
	GetAllUsers() ([]User, error)
	GetUsersPaginated(page int, limit int) ([]User, int64, error)
	CreateUser(user *User) error
	GetUserByID(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
	CreateUserWithTeacher(req CreateUserRequest) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *service) GetUsersPaginated(page, limit int) ([]User, int64, error) {
	offset := (page - 1) * limit
	return s.repo.GetUsersPaginated(offset, limit)
}

func (s *service) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) CreateUserWithTeacher(req CreateUserRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPassword

	teacher := &teachers.Teacher{
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		Email:            req.Email,
		Phone:            req.Phone,
		Department:       req.Department,
		ProfileImagePath: req.ProfileImagePath,
		Files:            req.Files,
	}

	user := &User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	return s.repo.CreateUserAndTeacher(user, teacher)
}

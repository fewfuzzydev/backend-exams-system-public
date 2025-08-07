package students

import "gorm.io/gorm"

type Repository interface {
	Create(student *Student) error
	FindAll() ([]Student, error)
	FindByID(id uint) (*Student, error)
	Update(student *Student) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Student{})
	return &repository{db}
}

func (r *repository) Create(student *Student) error {
	return r.db.Create(student).Error
}

func (r *repository) FindAll() ([]Student, error) {
	var students []Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *repository) FindByID(id uint) (*Student, error) {
	var student Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *repository) Update(student *Student) error {
	return r.db.Save(student).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Student{}, id).Error
}

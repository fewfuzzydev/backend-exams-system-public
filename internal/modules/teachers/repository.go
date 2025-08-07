package teachers

import "gorm.io/gorm"

type Repository interface {
	Create(user *Teacher) error
	FindAll() ([]Teacher, error)
	FindByID(id uint) (*Teacher, error)
	Update(user *Teacher) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Teacher{})
	return &repository{db}
}

func (r *repository) Create(teacher *Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *repository) FindAll() ([]Teacher, error) {
	var teachers []Teacher
	err := r.db.Find(&teachers).Error
	return teachers, err
}

func (r *repository) FindByID(id uint) (*Teacher, error) {
	var teacher Teacher
	err := r.db.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *repository) Update(teacher *Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Teacher{}, id).Error
}

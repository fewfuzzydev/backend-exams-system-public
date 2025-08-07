package exams

import "gorm.io/gorm"

type Repository interface {
	Create(user *Exams) error
	FindAll() ([]Exams, error)
	FindByID(id uint) (*Exams, error)
	Update(user *Exams) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Exams{})
	return &repository{db}
}

func (r *repository) Create(exams *Exams) error {
	return r.db.Create(exams).Error
}

func (r *repository) FindAll() ([]Exams, error) {
	var exams []Exams
	err := r.db.Find(&exams).Error
	return exams, err
}

func (r *repository) FindByID(id uint) (*Exams, error) {
	var exams Exams
	err := r.db.First(&exams, id).Error
	if err != nil {
		return nil, err
	}
	return &exams, nil
}

func (r *repository) Update(exams *Exams) error {
	return r.db.Save(exams).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Exams{}, id).Error
}

package examssubmission

import "gorm.io/gorm"

type Repository interface {
	Create(examSubmission *ExamSubmission) error
	FindAll() ([]ExamSubmission, error)
	FindByID(id uint) (*ExamSubmission, error)
	Update(examSubmission *ExamSubmission) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamSubmission{})
	return &repository{db}
}

func (r *repository) Create(examSubmission *ExamSubmission) error {
	return r.db.Create(examSubmission).Error
}

func (r *repository) FindAll() ([]ExamSubmission, error) {
	var examSubmissions []ExamSubmission
	err := r.db.Find(&examSubmissions).Error
	return examSubmissions, err
}

func (r *repository) FindByID(id uint) (*ExamSubmission, error) {
	var examSubmission ExamSubmission
	err := r.db.First(&examSubmission, id).Error
	if err != nil {
		return nil, err
	}
	return &examSubmission, nil
}

func (r *repository) Update(examSubmission *ExamSubmission) error {
	return r.db.Save(examSubmission).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ExamSubmission{}, id).Error
}

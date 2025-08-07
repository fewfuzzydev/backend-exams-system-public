package examssubmissionanswer

import "gorm.io/gorm"

type Repository interface {
	Create(examSubmissionAnswer *ExamSubmissionAnswer) error
	FindAll() ([]ExamSubmissionAnswer, error)
	FindByID(id uint) (*ExamSubmissionAnswer, error)
	Update(examSubmissionAnswer *ExamSubmissionAnswer) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamSubmissionAnswer{})
	return &repository{db}
}

func (r *repository) Create(examSubmissionAnswer *ExamSubmissionAnswer) error {
	return r.db.Create(examSubmissionAnswer).Error
}

func (r *repository) FindAll() ([]ExamSubmissionAnswer, error) {
	var examSubmissionAnswers []ExamSubmissionAnswer
	err := r.db.Find(&examSubmissionAnswers).Error
	return examSubmissionAnswers, err
}

func (r *repository) FindByID(id uint) (*ExamSubmissionAnswer, error) {
	var examSubmissionAnswer ExamSubmissionAnswer
	err := r.db.First(&examSubmissionAnswer, id).Error
	if err != nil {
		return nil, err
	}
	return &examSubmissionAnswer, nil
}

func (r *repository) Update(examSubmissionAnswer *ExamSubmissionAnswer) error {
	return r.db.Save(examSubmissionAnswer).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ExamSubmissionAnswer{}, id).Error
}

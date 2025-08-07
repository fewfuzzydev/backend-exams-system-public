package examsQuestions

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]ExamQuestion, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamQuestion{})
	return &repository{db}
}

func (r *repository) FindAll() ([]ExamQuestion, error) {
	var examsquestion []ExamQuestion
	err := r.db.Find(&examsquestion).Error
	return examsquestion, err
}

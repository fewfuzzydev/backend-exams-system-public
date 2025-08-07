package examschoice

import "gorm.io/gorm"

type Repository interface {
	Create(examChoice *ExamChoice) error
	FindAll() ([]ExamChoice, error)
	FindByID(id uint) (*ExamChoice, error)
	Update(examChoice *ExamChoice) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamChoice{})
	return &repository{db}
}

func (r *repository) Create(examChoice *ExamChoice) error {
	return r.db.Create(examChoice).Error
}

func (r *repository) FindAll() ([]ExamChoice, error) {
	var examChoices []ExamChoice
	err := r.db.Find(&examChoices).Error
	return examChoices, err
}

func (r *repository) FindByID(id uint) (*ExamChoice, error) {
	var examChoice ExamChoice
	err := r.db.First(&examChoice, id).Error
	if err != nil {
		return nil, err
	}
	return &examChoice, nil
}

func (r *repository) Update(examChoice *ExamChoice) error {
	return r.db.Save(examChoice).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ExamChoice{}, id).Error
}

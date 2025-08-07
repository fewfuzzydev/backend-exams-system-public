package examssessionexam

import "gorm.io/gorm"

type Repository interface {
	Create(examSessionExam *ExamSessionExam) error
	FindAll() ([]ExamSessionExam, error)
	FindByID(id uint) (*ExamSessionExam, error)
	Update(examSessionExam *ExamSessionExam) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamSessionExam{})
	return &repository{db}
}

func (r *repository) Create(examSessionExam *ExamSessionExam) error {
	return r.db.Create(examSessionExam).Error
}

func (r *repository) FindAll() ([]ExamSessionExam, error) {
	var examSessionExams []ExamSessionExam
	err := r.db.Find(&examSessionExams).Error
	return examSessionExams, err
}

func (r *repository) FindByID(id uint) (*ExamSessionExam, error) {
	var examSessionExam ExamSessionExam
	err := r.db.First(&examSessionExam, id).Error
	if err != nil {
		return nil, err
	}
	return &examSessionExam, nil
}

func (r *repository) Update(examSessionExam *ExamSessionExam) error {
	return r.db.Save(examSessionExam).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ExamSessionExam{}, id).Error
}

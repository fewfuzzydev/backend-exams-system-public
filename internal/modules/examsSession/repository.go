package examssession

import "gorm.io/gorm"

type Repository interface {
	Create(examSession *ExamSession) error
	FindAll() ([]ExamSession, error)
	FindByID(id uint) (*ExamSession, error)
	Update(examSession *ExamSession) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamSession{})
	return &repository{db}
}

func (r *repository) Create(examSession *ExamSession) error {
	return r.db.Create(examSession).Error
}

func (r *repository) FindAll() ([]ExamSession, error) {
	var examSessions []ExamSession
	err := r.db.Find(&examSessions).Error
	return examSessions, err
}

func (r *repository) FindByID(id uint) (*ExamSession, error) {
	var examSession ExamSession
	err := r.db.First(&examSession, id).Error
	if err != nil {
		return nil, err
	}
	return &examSession, nil
}

func (r *repository) Update(examSession *ExamSession) error {
	return r.db.Save(examSession).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ExamSession{}, id).Error
}

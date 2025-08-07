package subjects

import "gorm.io/gorm"

type Repository interface {
	Create(subject *Subject) error
	FindAll() ([]Subject, error)
	FindByID(id uint) (*Subject, error)
	Update(subject *Subject) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Subject{})
	return &repository{db}
}

func (r *repository) Create(subject *Subject) error {
	return r.db.Create(subject).Error
}

func (r *repository) FindAll() ([]Subject, error) {
	var subjects []Subject
	err := r.db.Find(&subjects).Error
	return subjects, err
}

func (r *repository) FindByID(id uint) (*Subject, error) {
	var subject Subject
	err := r.db.First(&subject, id).Error
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

func (r *repository) Update(subject *Subject) error {
	return r.db.Save(subject).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Subject{}, id).Error
}

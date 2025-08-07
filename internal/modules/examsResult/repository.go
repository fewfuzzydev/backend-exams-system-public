package examsResult

import "gorm.io/gorm"

type Repository interface {
	Create(data *ExamsResult) error
	FindAll() ([]ExamsResult, error)
	FindByID(id uint) (*ExamsResult, error)
	Update(data *ExamsResult) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&ExamsResult{})
	return &repository{db}
}

func (r *repository) Create(data *ExamsResult) error {
	return r.db.Create(data).Error
}

func (r *repository) FindAll() ([]ExamsResult, error) {
	var data []ExamsResult
	err := r.db.Find(&data).Error
	return data, err
}

func (r *repository) FindByID(id uint) (*ExamsResult, error) {
	var data ExamsResult
	err := r.db.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) Update(data *ExamsResult) error {
	return r.db.Save(data).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&ExamsResult{}, id).Error
}

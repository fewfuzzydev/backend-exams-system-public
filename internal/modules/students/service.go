package students

type Service interface {
	GetAll() ([]Student, error)
	Create(student *Student) error
	GetByID(id uint) (*Student, error)
	Update(student *Student) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]Student, error) {
	return s.repo.FindAll()
}

func (s *service) Create(student *Student) error {
	return s.repo.Create(student)
}

func (s *service) GetByID(id uint) (*Student, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(student *Student) error {
	return s.repo.Update(student)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

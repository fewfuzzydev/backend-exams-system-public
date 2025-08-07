package exams

type Service interface {
	GetAll() ([]Exams, error)
	Create(user *Exams) error
	GetByID(id uint) (*Exams, error)
	Update(user *Exams) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]Exams, error) {
	return s.repo.FindAll()
}

func (s *service) Create(user *Exams) error {
	return s.repo.Create(user)
}

func (s *service) GetByID(id uint) (*Exams, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(user *Exams) error {
	return s.repo.Update(user)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

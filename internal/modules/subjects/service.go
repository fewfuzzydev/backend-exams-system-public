package subjects

type Service interface {
	GetAll() ([]Subject, error)
	Create(subject *Subject) error
	GetByID(id uint) (*Subject, error)
	Update(subject *Subject) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]Subject, error) {
	return s.repo.FindAll()
}

func (s *service) Create(subject *Subject) error {
	return s.repo.Create(subject)
}

func (s *service) GetByID(id uint) (*Subject, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(subject *Subject) error {
	return s.repo.Update(subject)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

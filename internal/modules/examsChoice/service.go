package examschoice

type Service interface {
	GetAll() ([]ExamChoice, error)
	Create(examChoice *ExamChoice) error
	GetByID(id uint) (*ExamChoice, error)
	Update(examChoice *ExamChoice) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamChoice, error) {
	return s.repo.FindAll()
}

func (s *service) Create(examChoice *ExamChoice) error {
	return s.repo.Create(examChoice)
}

func (s *service) GetByID(id uint) (*ExamChoice, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(examChoice *ExamChoice) error {
	return s.repo.Update(examChoice)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

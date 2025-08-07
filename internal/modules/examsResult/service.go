package examsResult

type Service interface {
	GetAll() ([]ExamsResult, error)
	Create(data *ExamsResult) error
	GetByID(id uint) (*ExamsResult, error)
	Update(data *ExamsResult) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamsResult, error) {
	return s.repo.FindAll()
}

func (s *service) Create(data *ExamsResult) error {
	return s.repo.Create(data)
}

func (s *service) GetByID(id uint) (*ExamsResult, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(data *ExamsResult) error {
	return s.repo.Update(data)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

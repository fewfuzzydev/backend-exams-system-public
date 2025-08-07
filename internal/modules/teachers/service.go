package teachers

type Service interface {
	GetAll() ([]Teacher, error)
	Create(user *Teacher) error
	GetByID(id uint) (*Teacher, error)
	Update(user *Teacher) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]Teacher, error) {
	return s.repo.FindAll()
}

func (s *service) Create(user *Teacher) error {
	return s.repo.Create(user)
}

func (s *service) GetByID(id uint) (*Teacher, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(user *Teacher) error {
	return s.repo.Update(user)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

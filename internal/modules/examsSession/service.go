package examssession

type Service interface {
	GetAll() ([]ExamSession, error)
	Create(examSession *ExamSession) error
	GetByID(id uint) (*ExamSession, error)
	Update(examSession *ExamSession) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamSession, error) {
	return s.repo.FindAll()
}

func (s *service) Create(examSession *ExamSession) error {
	return s.repo.Create(examSession)
}

func (s *service) GetByID(id uint) (*ExamSession, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(examSession *ExamSession) error {
	return s.repo.Update(examSession)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

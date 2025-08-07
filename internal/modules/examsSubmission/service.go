package examssubmission

type Service interface {
	GetAll() ([]ExamSubmission, error)
	Create(examSubmission *ExamSubmission) error
	GetByID(id uint) (*ExamSubmission, error)
	Update(examSubmission *ExamSubmission) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamSubmission, error) {
	return s.repo.FindAll()
}

func (s *service) Create(examSubmission *ExamSubmission) error {
	return s.repo.Create(examSubmission)
}

func (s *service) GetByID(id uint) (*ExamSubmission, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(examSubmission *ExamSubmission) error {
	return s.repo.Update(examSubmission)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

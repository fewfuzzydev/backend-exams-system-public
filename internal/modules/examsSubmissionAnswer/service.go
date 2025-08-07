package examssubmissionanswer

type Service interface {
	GetAll() ([]ExamSubmissionAnswer, error)
	Create(examSubmissionAnswer *ExamSubmissionAnswer) error
	GetByID(id uint) (*ExamSubmissionAnswer, error)
	Update(examSubmissionAnswer *ExamSubmissionAnswer) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamSubmissionAnswer, error) {
	return s.repo.FindAll()
}

func (s *service) Create(examSubmissionAnswer *ExamSubmissionAnswer) error {
	return s.repo.Create(examSubmissionAnswer)
}

func (s *service) GetByID(id uint) (*ExamSubmissionAnswer, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(examSubmissionAnswer *ExamSubmissionAnswer) error {
	return s.repo.Update(examSubmissionAnswer)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

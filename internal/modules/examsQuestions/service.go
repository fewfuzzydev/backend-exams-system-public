package examsQuestions

type Service interface {
	GetAll() ([]ExamQuestion, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamQuestion, error) {
	return s.repo.FindAll()
}

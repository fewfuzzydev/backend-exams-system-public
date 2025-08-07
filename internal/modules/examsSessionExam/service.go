package examssessionexam

type Service interface {
	GetAll() ([]ExamSessionExam, error)
	Create(examSessionExam *ExamSessionExam) error
	GetByID(id uint) (*ExamSessionExam, error)
	Update(examSessionExam *ExamSessionExam) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]ExamSessionExam, error) {
	return s.repo.FindAll()
}

func (s *service) Create(examSessionExam *ExamSessionExam) error {
	return s.repo.Create(examSessionExam)
}

func (s *service) GetByID(id uint) (*ExamSessionExam, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(examSessionExam *ExamSessionExam) error {
	return s.repo.Update(examSessionExam)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

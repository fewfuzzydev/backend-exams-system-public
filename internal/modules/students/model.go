package students

import (
	"exams/internal/common"
	examssubmission "exams/internal/modules/examsSubmission"
)

type Student struct {
	Firstname string `gorm:"column:firstname"`
	Lastname  string `gorm:"column:lastname"`
	CardID    string `gorm:"column:card_id"`

	ExamSubmissions []examssubmission.ExamSubmission `gorm:"foreignKey:StudentID"`
	common.Audit
}

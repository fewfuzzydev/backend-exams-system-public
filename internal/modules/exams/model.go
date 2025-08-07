package exams

import (
	"exams/internal/common"
	"exams/internal/modules/examsQuestions"
)

type Exams struct {
	TeacherID   uint
	Title       string
	Description string
	SubjectID   uint

	Questions []examsQuestions.ExamQuestion `gorm:"foreignKey:ExamID"`
	common.Audit
}

package examschoice

import "exams/internal/common"

type ExamChoice struct {
	ID             uint `gorm:"primaryKey;column:examChoiceID"`
	ExamQuestionID uint
	ChoiceText     string
	IsCorrect      bool
	common.Audit
}

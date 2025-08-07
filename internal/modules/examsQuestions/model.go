package examsQuestions

import (
	"exams/internal/common"
	examschoice "exams/internal/modules/examsChoice"
)

type ExamQuestion struct {
	ExamID       uint   `gorm:"column:exam_id"`
	QuestionText string `gorm:"column:question_text"`
	QuestionType string `gorm:"column:question_type"`

	common.Audit

	Choices []examschoice.ExamChoice `gorm:"foreignKey:ExamQuestionID"`
}

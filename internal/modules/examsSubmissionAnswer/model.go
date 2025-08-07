package examssubmissionanswer

import (
	"exams/internal/common"
	"time"
)

type ExamSubmissionAnswer struct {
	ExamSubmissionsID uint
	ExamQuestionID    uint
	SelectedChoiceID  *uint
	StudentAnswer     string
	IsCorrect         bool
	Score             float64
	ReviewedAt        *time.Time
	ReviewedBy        *uint
	common.Audit
}

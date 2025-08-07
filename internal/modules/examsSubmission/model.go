package examssubmission

import (
	"exams/internal/common"
	examssubmissionanswer "exams/internal/modules/examsSubmissionAnswer"
	"time"
)

type ExamSubmission struct {
	StudentID      uint
	ExamSessionsID uint
	StartedAt      time.Time
	SubmittedAt    *time.Time
	TotalScore     float64

	Answers []examssubmissionanswer.ExamSubmissionAnswer `gorm:"foreignKey:ExamSubmissionsID"`
	common.Audit
}

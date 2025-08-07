package examssession

import (
	"exams/internal/common"
	examssessionexam "exams/internal/modules/examsSessionExam"
	examssubmission "exams/internal/modules/examsSubmission"
	"time"
)

type ExamSession struct {
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time

	ExamSessionExams []examssessionexam.ExamSessionExam `gorm:"foreignKey:ExamSessionsID"`
	ExamSubmissions  []examssubmission.ExamSubmission   `gorm:"foreignKey:ExamSessionsID"`
	common.Audit
}

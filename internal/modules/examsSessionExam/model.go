package examssessionexam

import "exams/internal/common"

type ExamSessionExam struct {
	ExamSessionsID uint
	ExamID         uint
	QuestionLimit  uint
	common.Audit
}

package subjects

import (
	"exams/internal/common"
	"exams/internal/modules/exams"
)

type Subject struct {
	SubjectName string `gorm:"column:subject_name"`
	Description string `gorm:"column:description"`

	Exams []exams.Exams `gorm:"foreignKey:SubjectID"`
	common.Audit
}

package teachers

import (
	"exams/internal/common"

	"github.com/lib/pq"
)

type Teacher struct {
	Firstname        string         `gorm:"column:firstname" json:"firstname"`
	Lastname         string         `gorm:"column:lastname" json:"lastname"`
	Email            string         `gorm:"column:email" json:"email"`
	Phone            string         `gorm:"column:phone" json:"phone"`
	Department       string         `gorm:"column:department" json:"department"`
	ProfileImagePath string         `gorm:"column:profile_image_path" json:"profileImagePath"`
	Files            pq.StringArray `gorm:"type:text[];column:files"`
	common.Audit
}

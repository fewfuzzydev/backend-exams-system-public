package users

import (
	"exams/internal/common"
	"exams/internal/modules/teachers"
)

type User struct {
	Username string           `gorm:"column:username;not null" json:"username"`
	Password string           `gorm:"column:password;not null" json:"-"`
	Role     string           `gorm:"column:role" json:"role"`
	Teacher  teachers.Teacher `gorm:"foreignKey:ID;references:ID"`
	common.Audit
}

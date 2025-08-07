package common

import (
	"gorm.io/gorm"
)

//	type Model struct {
//		ID        uint      `gorm:"primarykey" json:"id"`
//		CreatedAt time.Time `json:"createdAt"`
//		UpdatedAt time.Time `json:"updatedAt"`
//		DeletedAt time.Time `gorm:"index" json:"deletedAt"`
//	}
type Audit struct {
	CreatedBy uint `gorm:"column:createBy" json:"createdBy"`
	UpdatedBy uint `gorm:"column:updatedBy" json:"updatedBy"`
	DeletedBy uint `gorm:"column:deletedBy" json:"deletedBy"`
	gorm.Model
}

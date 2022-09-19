package global

import (
	"time"

	"gorm.io/gorm"
)

type GS_BASE_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id,string" form:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt" form:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" form:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                        // 删除时间
}

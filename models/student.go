package models

import (
	"time"

	"gorm.io/gorm"
)

// 创建本地数据库，并创建一张Student表，有字段：id（主键），student_no，name，gender，birth，往数据库中随机生产10万条学生信息
type Student struct {
	gorm.Model
	StudentNo int       `gorm:"column:student_no;index" json:"student_no" form:"student_no" binding:"required" unique:"true"`
	Name      string    `gorm:"column:name" json:"name" form:"name" binding:"required"`
	Gender    string    `gorm:"column:gender" json:"gender" form:"gender" binding:"required" enum:"男,女"`
	Birth     time.Time `gorm:"column:birth" json:"birth" form:"birth" binding:"required" time_format:"2006-01-02"`
}

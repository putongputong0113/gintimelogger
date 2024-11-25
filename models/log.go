// models/log.go

package models

import (
	"time"

	"gorm.io/gorm"
)

// RequestLog 存储请求的日志，包括请求路径、方法和耗时
type RequestLog struct {
	ID        uint    `gorm:"primaryKey"`
	Path      string  `gorm:"not null"`
	Method    string  `gorm:"not null"`
	Duration  float64 `gorm:"not null"` // 单位秒
	CreatedAt string
}

// CreateRequestLog 创建请求日志并保存到数据库
func CreateRequestLog(db *gorm.DB, path, method string, duration float64) error {
	log := RequestLog{
		Path:      path,
		Method:    method,
		Duration:  duration,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	return db.Create(&log).Error
}

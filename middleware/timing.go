// middleware/timing.go

package middleware

import (
	"fmt"
	"time"

	"github.com/putongputong0113/gintimelogger/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TimingMiddleware 用于记录请求的耗时
func TimingMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 执行请求
		c.Next()

		// 计算耗时
		duration := time.Since(start).Seconds()

		// 记录到数据库
		path := c.Request.URL.Path
		method := c.Request.Method
		err := models.CreateRequestLog(db, path, method, duration)
		if err != nil {
			// 如果记录日志失败，打印错误
			fmt.Printf("Error logging request: %v\n", err)
		}
	}
}

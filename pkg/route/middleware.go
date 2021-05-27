package route

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Limit 请求限制
func Limit(b int) gin.HandlerFunc {
	var (
		bucket = make(map[string]*rate.Limiter)
		lock   sync.Mutex
	)
	return func(c *gin.Context) {
		ip := c.ClientIP()

		lock.Lock()
		l, exist := bucket[ip]
		if !exist {
			l = rate.NewLimiter(rate.Limit(1), b)
			bucket[ip] = l
		}
		lock.Unlock()

		if l.Allow() {
			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

// LogFormat 日志格式化
func LogFormat(p gin.LogFormatterParams) string {
	return fmt.Sprintf("%s %s %d %s %d %s (%s)\n%s",
		p.TimeStamp.Format("2006/01/02 15:04:05"), p.Method, p.StatusCode,
		p.Path, p.BodySize, p.ClientIP, p.Latency, p.ErrorMessage,
	)
}

package hook

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"

	"github.com/dijkvy/geocentric/g2gorm"
	"github.com/dijkvy/geocentric/tag"
)

func NewTraceHook(levels []logger.LogLevel) g2gorm.Hook {
	return &traceHook{levels: levels}
}

// traceHook 主要用于将 context 中的 trace 信息注入到 gorm 产生的 logger 中
type traceHook struct {
	levels []logger.LogLevel
}

func (c *traceHook) Fire(e g2gorm.Entry) (err error) {
	// 提取 trace 信息并注入到日志中
	if value := tag.Extract(e.Context()); value != nil {
		_ = e.AppendField(zap.String(tag.Tr, fmt.Sprintf("%v", value)))
	}
	return nil
}

func (c *traceHook) Levels() (levels []logger.LogLevel) {
	return c.levels
}

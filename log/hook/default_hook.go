package hook

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dijkvy/geocentric/log"
	"github.com/dijkvy/geocentric/tag"
)

// NewTraceHook levels 日志级别
func NewTraceHook(levels []zapcore.Level) log.Hook {
	return &traceHook{levels: levels}
}

type traceHook struct {
	levels []zapcore.Level
}

func (c *traceHook) Levels() (lvs []zapcore.Level) {
	return c.levels
}

// Fire 从 context 上下文中提取信息
func (c *traceHook) Fire(e log.Entry) (err error) {
	if value := tag.Extract(e.Context()); value != nil {
		_ = e.AppendField(zap.String(tag.Tr, fmt.Sprintf("%v", value)))
	}
	return nil
}

package g2gorm

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

func NewWriter(logger *zap.Logger, optHooks ...Hook) (w Writer) {
	var hooks = make(Hooks, 0)
	for _, h := range optHooks {
		for _, level := range h.Levels() {
			hooks[level] = append(hooks[level], h)
		}
	}
	return &writerImpl{logger: logger, hooks: hooks}
}

type writerImpl struct {
	logger *zap.Logger
	hooks  Hooks
}

func (c *writerImpl) Printf(ctx context.Context, level logger.LogLevel, format string, values ...interface{}) {
	var e = newEntry(ctx)
	c.hooks.Fire(level, e)
	c.logger.Info(fmt.Sprintf(format, values...), e.GetFields()...)
}

package g2gorm

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func NewWriter(logger *zap.Logger, hook ...Hook) (w Writer) {
	return &writerImpl{logger: logger, hook: hook}
}

type writerImpl struct {
	logger *zap.Logger
	hook   []Hook
}

func (c *writerImpl) Printf(ctx context.Context, format string, values ...interface{}) {
	var e = newEntry(ctx)
	for _, h := range c.hook {

		if err := h.Fire(e); err != nil {
			fmt.Fprintf(os.Stderr, "failed to fire  %s", err)
			break
		}
	}
	c.logger.Info(fmt.Sprintf(format, values...), e.GetFields()...)
}

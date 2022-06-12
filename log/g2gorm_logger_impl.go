package log

import (
	"fmt"

	"go.uber.org/zap"
)

func NewG2LoggerHelper(logger *zap.Logger) (helper G2LoggerHelper) {
	return &g2gormLoggerImpl{logger: logger}
}

type g2gormLoggerImpl struct {
	logger *zap.Logger
}

func (c *g2gormLoggerImpl) Printf(format string, args ...interface{}) {
	c.logger.Info(fmt.Sprintf(format, args...))
}

package g2gorm

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

func NewLoggerInterface(logger *zap.Logger) logger.Interface {
	return &g2LoggerInterfaceImpl{logger: logger}
}

type g2LoggerInterfaceImpl struct {
	logger *zap.Logger
}

func (c *g2LoggerInterfaceImpl) LogMode(level logger.LogLevel) logger.Interface {
	panic("implement me")
}

func (c *g2LoggerInterfaceImpl) Info(ctx context.Context, s string, i ...interface{}) {
	panic("implement me")
}

func (c *g2LoggerInterfaceImpl) Warn(ctx context.Context, s string, i ...interface{}) {
	panic("implement me")
}

func (c *g2LoggerInterfaceImpl) Error(ctx context.Context, s string, i ...interface{}) {
	panic("implement me")
}

func (c *g2LoggerInterfaceImpl) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	panic("implement me")
}

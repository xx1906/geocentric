package log

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapFieldLogger struct {
	entry  Entry
	helper *zapLoggerHelper
	ctx    context.Context
}

func (c *zapFieldLogger) Debug(msg string, fields ...zap.Field) {
	if c.helper.levelEnabled(DebugLevel) {
		c.write(DebugLevel, msg, fields...)
	}
}

func (c *zapFieldLogger) Info(msg string, fields ...zap.Field) {
	if c.helper.levelEnabled(InfoLevel) {
		c.write(InfoLevel, msg, fields...)
	}
}

func (c *zapFieldLogger) Warn(msg string, fields ...zap.Field) {
	if c.helper.levelEnabled(WarnLevel) {
		c.write(WarnLevel, msg, fields...)
	}
}

func (c *zapFieldLogger) Error(msg string, fields ...zap.Field) {
	if c.helper.levelEnabled(ErrorLevel) {
		c.write(ErrorLevel, msg, fields...)
	}
}

func (c *zapFieldLogger) Fatal(msg string, fields ...zap.Field) {
	if c.helper.levelEnabled(FatalLevel) {
		c.write(FatalLevel, msg, fields...)
	}
}

// fireHooks execute hook, if get error print error
// msg to the StdOut
func (c *zapFieldLogger) fireHooks() {
	var tmpHooks = make(LevelHooks, len(c.helper.Hooks))
	for k, v := range c.helper.Hooks {
		tmpHooks[k] = v
	}
	if err := tmpHooks.Fire(c.helper.level, c.entry); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
	}
}

func (c *zapFieldLogger) write(level zapcore.Level, msg string, fields ...zap.Field) {
	c.fireHooks()
	if ce := c.helper.logger.Check(level, msg); ce != nil {
		ce.Write(append(fields, c.entry.GetFields()...)...)
	}
}

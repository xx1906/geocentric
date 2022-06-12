package log

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"
)

type zapFieldLogger struct {
	entry  Entry
	helper *zapLoggerHelper
	ctx    context.Context
}

func (c *zapFieldLogger) Debugf(format string, args ...interface{}) {
	if c.helper.levelEnabled(DebugLevel) {
		c.writef(DebugLevel, format, args...)
	}
}

func (c *zapFieldLogger) Infof(format string, args ...interface{}) {
	if c.helper.levelEnabled(InfoLevel) {
		c.writef(InfoLevel, format, args...)
	}
}

func (c *zapFieldLogger) Warnf(format string, args ...interface{}) {
	if c.helper.levelEnabled(WarnLevel) {
		c.writef(WarnLevel, format, args...)
	}
}

func (c *zapFieldLogger) Errorf(format string, args ...interface{}) {
	if c.helper.levelEnabled(ErrorLevel) {
		c.writef(ErrorLevel, format, args...)
	}
}

func (c *zapFieldLogger) Fatalf(format string, args ...interface{}) {
	if c.helper.levelEnabled(FatalLevel) {
		c.writef(FatalLevel, format, args...)
	}
}

func (c *zapFieldLogger) Debug(args ...interface{}) {
	if c.helper.levelEnabled(DebugLevel) {
		c.write(DebugLevel, args...)
	}
}

func (c *zapFieldLogger) Info(args ...interface{}) {
	if c.helper.levelEnabled(InfoLevel) {
		c.write(InfoLevel, args...)
	}
}

func (c *zapFieldLogger) Warn(args ...interface{}) {
	if c.helper.levelEnabled(WarnLevel) {
		c.write(WarnLevel, args...)
	}
}

func (c *zapFieldLogger) Error(args ...interface{}) {
	if c.helper.levelEnabled(ErrorLevel) {
		c.write(ErrorLevel, args...)
	}
}

func (c *zapFieldLogger) Fatal(args ...interface{}) {
	if c.helper.levelEnabled(FatalLevel) {
		c.write(FatalLevel, args...)
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

func (c *zapFieldLogger) pollDataf(format string, args ...interface{}) (data string) {
	buff, release := c.helper.pollBuff()
	defer release()
	buff.WriteString(fmt.Sprintf(format, args...))
	return buff.String()
}

func (c *zapFieldLogger) pollData(args ...interface{}) (data string) {
	buff, release := c.helper.pollBuff()
	defer release()
	for _, v := range args {
		buff.WriteString(fmt.Sprintf("%v ", v))
	}
	return buff.String()
}

func (c *zapFieldLogger) write(level zapcore.Level, args ...interface{}) {
	msg := c.pollData(args...)
	c.fireHooks()
	if ce := c.helper.logger.Check(level, msg); ce != nil {
		ce.Write(c.entry.GetFields()...)
	}
}

func (c *zapFieldLogger) writef(level zapcore.Level, format string, args ...interface{}) {
	msg := c.pollDataf(format, args...)
	c.fireHooks()
	if ce := c.helper.logger.Check(level, msg); ce != nil {
		ce.Write(c.entry.GetFields()...)
	}
}

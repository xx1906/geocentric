package g2gorm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Colors
const (
	Reset       = logger.Reset
	Red         = logger.Red
	Green       = logger.Green
	Yellow      = logger.Yellow
	Blue        = logger.Blue
	Magenta     = logger.Magenta
	Cyan        = logger.Cyan
	White       = logger.White
	BlueBold    = logger.BlueBold
	MagentaBold = logger.MagentaBold
	RedBold     = logger.RedBold
	YellowBold  = logger.YellowBold
)

const (
	// Silent silent log level
	Silent = logger.Silent
	// Info info log level
	Info = logger.Info
	// Warn warn log level
	Warn = logger.Warn
	// Error error log level
	Error = logger.Error
)

func NewLoggerInterface(writer Writer, config logger.Config) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr = BlueBold + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	}

	return &g2LoggerInterfaceImpl{
		Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type g2LoggerInterfaceImpl struct {
	logger.Config
	Writer
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

type Writer interface {
	Printf(ctx context.Context, format string, values ...interface{})
}

func (c *g2LoggerInterfaceImpl) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *c
	newlogger.LogLevel = level
	return &newlogger
}

func (c *g2LoggerInterfaceImpl) Info(ctx context.Context, format string, data ...interface{}) {
	if c.LogLevel >= Info {
		c.Printf(ctx, c.infoStr+format, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (c *g2LoggerInterfaceImpl) Warn(ctx context.Context, format string, data ...interface{}) {
	if c.LogLevel >= Warn {
		c.Printf(ctx, c.warnStr+format, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (c *g2LoggerInterfaceImpl) Error(ctx context.Context, format string, data ...interface{}) {
	if c.LogLevel >= Error {
		c.Printf(ctx, c.errStr+format, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (c *g2LoggerInterfaceImpl) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if c.LogLevel <= Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && c.LogLevel >= Error && (!errors.Is(err, logger.ErrRecordNotFound) || !c.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			c.Printf(ctx, c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			c.Printf(ctx, c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > c.SlowThreshold && c.SlowThreshold != 0 && c.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", c.SlowThreshold)
		if rows == -1 {
			c.Printf(ctx, c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			c.Printf(ctx, c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case c.LogLevel == Info:
		sql, rows := fc()
		if rows == -1 {
			c.Printf(ctx, c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			c.Printf(ctx, c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

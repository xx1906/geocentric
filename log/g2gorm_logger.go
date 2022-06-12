package log

// G2LoggerHelper gorm 日志接口
type G2LoggerHelper interface {
	Printf(format string, args ...interface{})
}

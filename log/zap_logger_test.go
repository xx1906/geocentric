package log

import (
	"testing"

	"go.uber.org/zap"

	"github.com/djikvy/geocentric/log/config"
)

func TestNewZapLogger(t *testing.T) {

	logger, err := NewZapLogger(&config.ZapConfig{Path: "."},
		zap.WithCaller(true),
		zap.AddCallerSkip(2))
	if err != nil {
		t.Error(err)
		return
	}
	logger.Core()
}

func Test_getWriter(t *testing.T) {
	writer := getWriter("app.log", 1, 1, 1, false)
	if writer == nil {
		t.Error("expect not nil")
	}
}

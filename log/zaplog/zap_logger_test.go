package zaplog

import (
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/dijkvy/geocentric/log/config"
)

func TestNewZapLogger(t *testing.T) {

	logger, err := NewZapLogger(&config.ZapConfig{Path: "."})
	if err != nil {
		t.Error(err)
		return
	}
	logger.WithOptions(zap.WithCaller(true), zap.AddCallerSkip(2))
	defer logger.Sync()
	logger.Warn("any")
	logger.Info("", zap.String("", "value"), zap.Any("hello", 233))

}

func Test_getWriter(t *testing.T) {
	fileName := "app.log"
	writer := getWriter(fileName, 1, 1, 1, false)
	if writer == nil {
		t.Error("expect not nil")
		return
	}
	defer os.Remove(fileName)
	writer.Write([]byte("hello world"))
}

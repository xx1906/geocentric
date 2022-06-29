package zaplog

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dijkvy/geocentric/config"
)

func TestNewZapLogger(t *testing.T) {

	logger, err := NewZapLogger(&config.ZapConfig{FileName: "biz.log"})
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

func Test_parseLogLevel(t *testing.T) {
	type testTab struct {
		inputStr string
		expected zapcore.Level
	}
	// 测试表格
	var tables = []testTab{
		testTab{inputStr: LevelDebug, expected: zapcore.DebugLevel},
		testTab{inputStr: LevelInfo, expected: zapcore.InfoLevel},
		testTab{inputStr: LevelWarn, expected: zapcore.WarnLevel},
		testTab{inputStr: LevelError, expected: zapcore.ErrorLevel},
		testTab{inputStr: LevelFatal, expected: zapcore.FatalLevel},
		testTab{inputStr: "_info", expected: zapcore.InfoLevel},
		testTab{inputStr: "''", expected: zapcore.InfoLevel},
	}

	for _, v := range tables {
		output, err := parseLogLevel(v.inputStr)
		if err != nil {
			t.Error(err)
			return
		}
		if output != v.expected {
			t.Errorf("expected %v but get %v", v.expected, output)
		}
		t.Log(output, v.expected, v.inputStr)
	}
}

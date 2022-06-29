package log

import (
	"context"
	stdLog "log"

	"go.uber.org/zap"

	"github.com/dijkvy/geocentric/config"
	"github.com/dijkvy/geocentric/log/zaplog"
)

func ExampleNewHelper() {
	var logger, err = zaplog.NewZapLogger(&config.ZapConfig{})
	if err != nil {
		stdLog.Fatal("err ", err)
	}
	var helper Helper = NewHelper(logger)
	helper.WithContext(context.Background()).Info("hello")
	helper.WithContext(context.TODO()).Error("error ", zap.String("error", "error info"))
}

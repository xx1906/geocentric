package log

import (
	"context"
	stdLog "log"

	"github.com/dijkvy/geocentric/log/config"
	"github.com/dijkvy/geocentric/log/zaplog"
)

func ExampleZapLoggerHelper_WithContext() {
	var logger, err = zaplog.NewZapLogger(&config.ZapConfig{})
	if err != nil {
		stdLog.Fatal("err ", err)
	}
	var helper Helper = NewZapHelper(logger)
	helper.WithContext(context.Background()).Info("hello")
	helper.WithContext(context.TODO()).Error("error ", "error info")
}

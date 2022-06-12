package log

import (
	"context"
	stdLog "log"

	"github.com/dijkvy/geocentric/log/config"
)

func ExampleZapLoggerHelper_WithContext() {
	var logger, err = NewZapLogger(&config.ZapConfig{})
	if err != nil {
		stdLog.Fatal("err ", err)
	}
	var helper Helper = NewZapHelper(logger)
	helper.WithContext(context.Background()).Info("hello")
	helper.WithContext(context.TODO()).Error("error ", "error info")
}

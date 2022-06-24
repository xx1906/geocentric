package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dijkvy/geocentric/log"
	"github.com/dijkvy/geocentric/log/config"
	"github.com/dijkvy/geocentric/log/zaplog"
)

var (
	cfgPath = flag.String("config", "log/examples/zaplog/configs/zaplog.toml", "zaplog config path")
)

type key int

func (c key) Levels() (lvs []zapcore.Level) {
	return []zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel, zapcore.PanicLevel}
}

func (c key) Fire(e log.Entry) (err error) {
	if e.Context().Value(_key) == nil {
		return err
	}
	return e.AppendField(zap.Any("tr", e.Context().Value(_key)))
}

var _key key

func main() {
	flag.Parse()
	var conf config.ZapConfig
	var err error
	fmt.Println(*cfgPath)

	if _, err = toml.DecodeFile(*cfgPath, &conf); err != nil {
		panic(err)
	}

	logger, err := zaplog.NewZapLogger(&conf)
	if err != nil {
		panic(err)
	}
	logger = logger.WithOptions(zap.WithCaller(true), zap.AddCallerSkip(2))
	defer logger.Sync()
	opts := make([]log.Option, 0)
	opts = append(opts, log.AddHook(&injectAppName{AppName: "demo"}), log.AddHook(key(0)))
	helper := log.NewZapHelper(logger, opts...)
	ctx := context.TODO()
	for i := 0; i < 1000; i++ {

		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Debug("debug")
		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Info("info")
		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Warn("warn")
		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Error("error")

		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Debug("debug  ", zap.Time("now", time.Now()))
		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Info("info ", zap.Time("now", time.Now()))
		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Warn("info ", zap.Time("now", time.Now()))
		helper.WithContext(context.WithValue(ctx, _key, uuid.New().String())).Error("error ", zap.Time("now", time.Now()))
	}
}

type injectAppName struct {
	AppName string
}

func (c *injectAppName) Levels() (lvs []zapcore.Level) {
	return []zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel, zapcore.PanicLevel}
}

func (c *injectAppName) Fire(e log.Entry) (err error) {
	_ = e.Context()
	e.AppendField(zap.String("app_name", c.AppName))
	return nil
}

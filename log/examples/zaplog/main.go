package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dijkvy/geocentric/log"
	"github.com/dijkvy/geocentric/log/config"
)

var (
	cfgPath = flag.String("config", "log/examples/zaplog/configs/zaplog.toml", "zaplog config path")
)

func main() {
	flag.Parse()
	var conf config.ZapConfig
	var err error
	fmt.Println(*cfgPath)

	if _, err = toml.DecodeFile(*cfgPath, &conf); err != nil {
		panic(err)
	}
	defer func() {
		abs, err := filepath.Abs(conf.GetPath())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(os.RemoveAll(abs))
	}()

	logger, err := log.NewZapLogger(&conf, zap.WithCaller(true), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	opts := make([]log.Option, 0)
	opts = append(opts, log.AddHook(&injectAppName{AppName: "demo"}))
	helper := log.NewZapHelper(logger, opts...)
	for i := 0; i < 100000000; i++ {

		helper.WithContext(context.TODO()).Debug("debug")
		helper.WithContext(context.TODO()).Info("info")
		helper.WithContext(context.TODO()).Warn("warn")
		helper.WithContext(context.TODO()).Error("error")

		helper.WithContext(context.TODO()).Debugf("debug  time:%s", time.Now())
		helper.WithContext(context.TODO()).Infof("info  time:%s", time.Now())
		helper.WithContext(context.TODO()).Warnf("info  time:%s", time.Now())
		helper.WithContext(context.TODO()).Errorf("error  time:%s", time.Now())
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

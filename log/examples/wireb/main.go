package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/dijkvy/geocentric/config"
)

var (
	cfgPath = flag.String("config", "log/examples/wireb/configs/zaplog.toml", "zaplog config path")
)

func main() {
	flag.Parse()
	var conf config.ZapConfig
	var err error
	fmt.Println(*cfgPath)

	if _, err = toml.DecodeFile(*cfgPath, &conf); err != nil {
		panic(err)
	}
	logger, err := InitZapLogger(&conf)
	if err != nil {
		panic(err)
	}
	helper, err := InitLoggerHelper(logger)
	if err != nil {
		panic(err)
	}
	helper.WithContext(context.TODO()).Info("this test case")
	helper.WithContext(context.TODO()).Warn("this test case")
	helper.WithContext(context.TODO()).Error("this test case")
}

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/dijkvy/geocentric/config"
	"github.com/dijkvy/geocentric/log"
	"github.com/dijkvy/geocentric/log/zaplog"
)

func InitLoggerHelper(logger *zap.Logger, opt ...log.Option) (helper log.Helper, err error) {
	panic(wire.Build(log.LogHelperProvider))
}

func InitZapLogger(cfg *config.ZapConfig) (logger *zap.Logger, err error) {
	panic(wire.Build(zaplog.ZapLogProvider))
}

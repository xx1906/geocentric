package zaplog

import "github.com/google/wire"

var ZapLogProvider = wire.NewSet(NewZapLogger)

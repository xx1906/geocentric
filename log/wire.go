package log

import "github.com/google/wire"

var LogHelperProvider = wire.NewSet(NewHelper)

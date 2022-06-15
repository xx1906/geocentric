package log

import (
	"context"
	"testing"

	"go.uber.org/zap"
)

func TestZapLoggerEntry_AppendField(t *testing.T) {
	var entry Entry = newZapLogEntry(context.TODO())
	entry.AppendField(zap.String("test", "test"))
	if len(entry.GetFields()) != 1 {
		t.Error("not ok")
	}
}

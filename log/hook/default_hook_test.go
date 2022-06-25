package hook

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dijkvy/geocentric/log"
	"github.com/dijkvy/geocentric/tag"
)

type myEntryImpl struct {
	ctx    context.Context
	fields []zap.Field
}

func (c *myEntryImpl) Context() (ctx context.Context) {
	return c.ctx
}

func (c *myEntryImpl) AppendField(field zap.Field) (err error) {
	c.fields = append(c.fields, field)
	return nil
}

func (c *myEntryImpl) GetFields() (fields []zap.Field) {
	return c.fields
}

func TestNewTraceHook(t *testing.T) {
	var levels = append(make([]zapcore.Level, 0),
		zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel, zapcore.FatalLevel)

	var h log.Hook = NewTraceHook(levels)
	if reflect.DeepEqual(h.Levels(), levels) {
		t.Log("pass")
	}
}

func Test_traceHook_Fire(t *testing.T) {
	var levels = append(make([]zapcore.Level, 0),
		zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel, zapcore.FatalLevel)

	var h log.Hook = NewTraceHook(levels)
	if reflect.DeepEqual(h.Levels(), levels) {
		t.Log("pass")
	}
	var entry log.Entry = &myEntryImpl{ctx: tag.Inject(context.TODO(), "trace_data"), fields: make([]zapcore.Field, 0)}
	_ = h.Fire(entry)
	if len(entry.GetFields()) != 1 {
		t.Error("field should be 1, but get ", entry.GetFields())

	}
	for _, v := range entry.GetFields() {
		t.Log(v.Key, v.String)
	}
}

func Test_traceHook_Levels(t *testing.T) {
	// pass
}

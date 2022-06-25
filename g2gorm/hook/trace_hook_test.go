package hook

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"

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
	var levels = append(make([]logger.LogLevel, 0), logger.Silent, logger.Info, logger.Warn, logger.Error)
	var h = NewTraceHook(levels)
	if reflect.DeepEqual(h.Levels(), levels) {
		t.Log("pass")
	}

}

func Test_traceHook_Fire(t *testing.T) {
	var levels = append(make([]logger.LogLevel, 0), logger.Silent, logger.Info, logger.Warn, logger.Error)
	var h = NewTraceHook(levels)
	var entry = myEntryImpl{ctx: context.TODO()}
	entry.ctx = tag.Inject(entry.ctx, "data")
	_ = h.Fire(&entry)
	if len(entry.GetFields()) != 1 {
		t.Error("len should be 1")
	}
	for _, v := range entry.GetFields() {
		t.Log(v.Key, v.String)
	}
}

func Test_traceHook_Levels(t *testing.T) {
	// pass
}

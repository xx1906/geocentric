package log

import (
	"context"

	"go.uber.org/zap"
)

type zapLoggerEntry struct {
	ctx    context.Context // context
	fields []zap.Field     // data gen by Entry impl
}

// Context get ctx
func (c *zapLoggerEntry) Context() (ctx context.Context) {
	return c.ctx
}

// AppendField append field to fields
func (c *zapLoggerEntry) AppendField(field zap.Field) (err error) {
	c.fields = append(c.fields, field)
	return nil
}

// GetFields get fields
func (c *zapLoggerEntry) GetFields() []zap.Field {
	return c.fields
}

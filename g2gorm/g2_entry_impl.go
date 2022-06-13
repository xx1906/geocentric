package g2gorm

import (
	"context"

	"go.uber.org/zap"
)

type entryImpl struct {
	ctx    context.Context
	fields []zap.Field
}

func (c *entryImpl) Context() (ctx context.Context) {
	return c.ctx
}

func (c *entryImpl) AppendField(field zap.Field) (err error) {
	c.fields = append(c.fields, field)
	return nil
}

func (c *entryImpl) GetFields() (fields []zap.Field) {
	return c.fields
}

func newEntry(ctx context.Context) (f Entry) {
	return &entryImpl{
		ctx:    ctx,
		fields: make([]zap.Field, 0),
	}
}

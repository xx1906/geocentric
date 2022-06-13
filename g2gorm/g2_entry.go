package g2gorm

import (
	"context"

	"go.uber.org/zap"
)

type Entry interface {
	Context() (ctx context.Context)
	AppendField(field zap.Field) (err error)
	GetFields() (fields []zap.Field)
}

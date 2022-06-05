package log

import (
	"context"
)

// Helper logger helper
type Helper interface {
	WithContext(ctx context.Context) (logger FieldLogger)
}

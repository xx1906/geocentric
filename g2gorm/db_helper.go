package g2gorm

import (
	"context"

	"gorm.io/gorm"
)

type DBHelper interface {
	WithContext(ctx context.Context) (db *gorm.DB)
}

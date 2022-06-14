package g2gorm

import (
	"fmt"
	"os"

	"gorm.io/gorm/logger"
)

type Hooks map[logger.LogLevel][]Hook
type Hook interface {
	Fire(e Entry) (err error)
	Levels() (levels []logger.LogLevel)
}

func (c Hooks) Fire(level logger.LogLevel, e Entry) {
	hooks := c[level]
	for _, h := range hooks {
		if err := h.Fire(e); err != nil {
			fmt.Fprintf(os.Stderr, "failed to fire %s", err)
			continue
		}
	}
}

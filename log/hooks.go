package log

import (
	"go.uber.org/zap/zapcore"
)

type LevelHooks map[zapcore.Level][]Hook

type Hook interface {
	Levels() (lvs []zapcore.Level)
	Fire(e Entry) (err error)
}

// Add hook plugin
func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

// Fire execute hook plugin, if get error return err else return nil
func (hooks LevelHooks) Fire(level zapcore.Level, entry Entry) (err error) {
	for _, hook := range hooks[level] {
		if err = hook.Fire(entry); err != nil {
			return err
		}
	}
	return nil
}

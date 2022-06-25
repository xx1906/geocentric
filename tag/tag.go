package tag

import "context"

type traceKey int

const Tr = "tr"

var (
	key traceKey = 1
)

// Inject 注入 value
func Inject(ctx context.Context, value interface{}) context.Context {
	// 如果存在, 则不需要注入
	if Extract(ctx) != nil {
		return ctx
	}
	return context.WithValue(ctx, key, value)
}

// Extract 提取 value
func Extract(ctx context.Context) (any interface{}) {
	return ctx.Value(key)
}

package tag

import (
	"context"
	"testing"
)

func TestInject(t *testing.T) {
	ctx := context.TODO()
	ctx = Inject(ctx, "inject_value")
	t.Log(ctx)
}

func TestExtract(t *testing.T) {
	ctx := context.TODO()
	if value := Extract(ctx); value == nil {
		t.Log("not value")
		ctx = Inject(ctx, "value")
	}

	if value := Extract(ctx); value == nil {
		t.Error("should have value but not")
	}
}

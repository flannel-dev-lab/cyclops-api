package cycapi

import (
	"context"
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	ctx := context.Background()

	ctx = Add(ctx, "hello", "world")
	ctx = Add(ctx, "go", "best language")

	Error(ctx, "test err", errors.New("testing"))
}

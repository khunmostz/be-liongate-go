package common

import (
	"context"
	"time"
)

// ContextWithTimeout adds a timeout to the provided context
func ContextWithTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, DefaultTimeout)
}

// DefaultTimeout is the default timeout for database operations
const DefaultTimeout = 10 * time.Second

package services

import (
	"context"
	"time"
)

type (
	PreProcess interface {
		PreStart(ctx context.Context, dir string) (string, time.Time, error)
	}
)

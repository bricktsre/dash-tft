package services

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/exp/slog"
)

func NewCount(log *slog.Logger) Count {
	return Count{
		Log: log
	}
}

type Count struct {
	Log *slog.Logger
}

func 
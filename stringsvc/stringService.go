package main

import (
	"context"
	"errors"
	"strings"
)

type (
	// StringService : Minimal go-kit service
	StringService interface {
		Uppercase(context.Context, string) (string, error)
		Count(context.Context, string) int
	}

	stringService struct{}
)

var (
	// ErrEmpty : the string is empty
	ErrEmpty = errors.New("Empty string")
)

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(s), nil
}

func (stringService) Count(_ context.Context, s string) int {
	return len(s)
}

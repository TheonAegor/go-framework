package config

import (
	"context"
	"errors"
)

var (
	// ErrInvalidStruct is returned when the target struct is invalid
	ErrInvalidStruct = errors.New("invalid struct specified")
	// ErrCodecMissing is returned when codec needed and not specified
	ErrCodecMissing = errors.New("codec missing")
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Name returns name of config
	Name() string
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Load config from sources
	Load(context.Context, ...LoadOption) error
	// String returns config type name
	String() string
	// Watch and Save ?
}

// Load loads config from config sources
func Load(ctx context.Context, cs ...Config) error {
	var err error
	for _, c := range cs {
		if err = c.Init(); err != nil {
			return err
		}
		if err = c.Load(ctx); err != nil {
			return err
		}
	}
	return nil
}

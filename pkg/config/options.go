package config

import (
	"context"
	"github.com/TheonAegor/go-framework/pkg/codec"
	"github.com/TheonAegor/go-framework/pkg/logger"
	"github.com/TheonAegor/go-framework/pkg/meter"
	"github.com/TheonAegor/go-framework/pkg/tracer"
)

// Options hold the config options
type Options struct {
	Name       string
	AllowFail  bool
	BeforeLoad []func(context.Context, Config) error
	AfterLoad  []func(context.Context, Config) error
	BeforeSave []func(context.Context, Config) error
	AfterSave  []func(context.Context, Config) error
	// Struct that holds config data
	Struct interface{}
	// StructTag name
	StructTag string
	// Logger that will be used
	Logger logger.Loggerer
	// Tracer used for trace
	Tracer tracer.Tracer
	// Meter
	Meter meter.Meter
	// Codec that used for load/save
	Codec codec.Codec
	// Context for alternative data
	Context context.Context
}

// Option function signature
type Option func(o *Options)

// NewOptions new options struct with filed values
func NewOptions(opts ...Option) Options {
	options := Options{
		Logger:  logger.DefaultLogger,
		Meter:   meter.DefaultMeter,
		Tracer:  tracer.DefaultTracer,
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}

	return options
}

// LoadOption function signature
type LoadOption func(o *LoadOptions)

// LoadOptions struct
type LoadOptions struct {
	Struct   interface{}
	Override bool
	Append   bool
	Context  context.Context
}

// NewLoadOptions create LoadOptions struct with provided opts
func NewLoadOptions(opts ...LoadOption) LoadOptions {
	options := LoadOptions{}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// LoadOverride override values when load
func LoadOverride(b bool) LoadOption {
	return func(o *LoadOptions) {
		o.Override = b
	}
}

// LoadAppend override values when load
func LoadAppend(b bool) LoadOption {
	return func(o *LoadOptions) {
		o.Append = b
	}
}

// LoadStruct override struct for loading
func LoadStruct(src interface{}) LoadOption {
	return func(o *LoadOptions) {
		o.Struct = src
	}
}

// SaveOption function signature
type SaveOption func(o *SaveOptions)

// SaveOptions struct
type SaveOptions struct {
	Struct  interface{}
	Context context.Context
}

// SaveStruct override struct for save to config
func SaveStruct(src interface{}) SaveOption {
	return func(o *SaveOptions) {
		o.Struct = src
	}
}

// NewSaveOptions fill SaveOptions struct
func NewSaveOptions(opts ...SaveOption) SaveOptions {
	options := SaveOptions{}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// AllowFail allows config source to fail
func AllowFail(b bool) Option {
	return func(o *Options) {
		o.AllowFail = b
	}
}

// BeforeLoad run funcs before config load
func BeforeLoad(fn ...func(context.Context, Config) error) Option {
	return func(o *Options) {
		o.BeforeLoad = fn
	}
}

// AfterLoad run funcs after config load
func AfterLoad(fn ...func(context.Context, Config) error) Option {
	return func(o *Options) {
		o.AfterLoad = fn
	}
}

// BeforeSave run funcs before save
func BeforeSave(fn ...func(context.Context, Config) error) Option {
	return func(o *Options) {
		o.BeforeSave = fn
	}
}

// AfterSave run fncs after save
func AfterSave(fn ...func(context.Context, Config) error) Option {
	return func(o *Options) {
		o.AfterSave = fn
	}
}

// Context pass context
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// Codec sets the source codec
func Codec(c codec.Codec) Option {
	return func(o *Options) {
		o.Codec = c
	}
}

// Logger sets the logger
func Logger(l logger.Loggerer) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// Tracer to be used for tracing
func Tracer(t tracer.Tracer) Option {
	return func(o *Options) {
		o.Tracer = t
	}
}

// Struct used as config
func Struct(v interface{}) Option {
	return func(o *Options) {
		o.Struct = v
	}
}

// StructTag sets the struct tag that used for filling
func StructTag(name string) Option {
	return func(o *Options) {
		o.StructTag = name
	}
}

// Name sets the name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

/**
 * @Package zapctx
 * @Time: 2022/2/15 20:28 PM
 * @Author: wuhb
 * @File: zapctx.go
 */

package zapctx

import (
	"context"
	"io"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(cfg *config)
type CtxDecoderFunc func(ctx context.Context, l *zap.Logger) *zap.Logger
type config struct {
	path       string
	rotate     *RotateConfig
	level      zapcore.Level
	caller     int
	debug      bool
	writer     io.Writer
	ctxDecoder CtxDecoderFunc
}
type Logger struct {
	*zap.Logger
	ctxDecoder CtxDecoderFunc
}

func (l *Logger) clone() *Logger {
	c := *l
	return &c
}
func (l *Logger) WithContext(ctx context.Context) *Logger {
	if l.ctxDecoder != nil {
		nl := l.clone()
		nl.Logger = l.ctxDecoder(ctx, nl.Logger)
		return nl
	}
	return l
}

func (l *Logger) Named(name string) *Logger {
	if name == "" {
		return l
	}
	nl := l.clone()
	nl.Logger = nl.Logger.Named(name)
	return nl
}

func (l *Logger) Zap() *zap.Logger {
	return l.Logger
}

func New(opts ...Option) *Logger {
	cfg := &config{
		rotate:     &RotateConfig{},
		level:      zap.DebugLevel,
		caller:     0,
		debug:      false,
		writer:     log.Writer(),
		ctxDecoder: nil,
	}
	cfg.rotate.Init()
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.path != "" {
		cfg.writer = Rotate(cfg.path, cfg.rotate)
	}
	return &Logger{
		Logger:     initZapLogger(cfg),
		ctxDecoder: cfg.ctxDecoder,
	}
}

func NewWithZap(logger *zap.Logger) *Logger {
	return &Logger{
		Logger:     logger,
		ctxDecoder: nil,
	}
}
func WithTraceIDOption(traceID string) Option {
	return func(cfg *config) {
		cfg.ctxDecoder = func(ctx context.Context, l *zap.Logger) *zap.Logger {
			return traceDecoder(ctx, traceID, l)
		}
	}
}

func WithDebugOption(debug bool) Option {
	return func(cfg *config) {
		cfg.debug = debug
	}
}

func WithLevelOption(l string) Option {
	return func(cfg *config) {
		cfg.level = convertToZapLevel(l)
	}
}

func WithCtxDecoderOption(decoder CtxDecoderFunc) Option {
	return func(cfg *config) {
		cfg.ctxDecoder = decoder
	}
}

func WithRotateOption(maxAge, maxBackups int, compress bool) Option {
	return func(cfg *config) {
		cfg.rotate.MaxAge = maxAge
		cfg.rotate.MaxBackups = maxBackups
		cfg.rotate.Compress = compress
	}
}

func WithLogPathOption(path string) Option {
	return func(cfg *config) {
		cfg.path = path
	}
}

func traceDecoder(ctx context.Context, traceID string, l *zap.Logger) *zap.Logger {
	trace := ctx.Value(traceID)
	if trace == nil {
		return l
	}
	l = l.With(zap.Any(traceID, trace))
	return l
}

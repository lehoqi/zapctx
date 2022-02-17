/**
 * @Package zapctx
 * @Time: 2022/2/15 20:28 PM
 * @Author: wuhb
 * @File: zap.go
 */

package zapctx

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZapLogger(cfg *config) *zap.Logger {
	var core zapcore.Core
	var encoderCfg zapcore.EncoderConfig
	var opts []zap.Option
	if cfg.debug {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	if cfg.path == "" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	core = zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(cfg.writer), cfg.level)
	opts = []zap.Option{zap.AddCaller()}
	if cfg.caller > 0 {
		opts = append(opts, zap.AddCallerSkip(cfg.caller))
	}
	return zap.New(core, opts...)
}

func ContextField(key string, ctx context.Context) zap.Field {
	val := ctx.Value(key)
	if val == nil {
		return zap.Skip()
	}
	return zap.Any(key, val)
}

/**
 * @Package zapctx
 * @Time: 2022/2/15 20:24 PM
 * @Author: wuhb
 * @File: level.go
 */

package zapctx

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

const (
	DEBUG  = "debug"
	INFO   = "info"
	WARN   = "warn"
	ERROR  = "error"
	FATAL  = "fatal"
	PANIC  = "panic"
	DPANIC = "dpanic"
)

func convertToZapLevel(l string) zapcore.Level {
	lowerL := strings.ToLower(l)
	switch lowerL {
	case DEBUG:
		return zapcore.DebugLevel
	case INFO:
		return zapcore.InfoLevel
	case WARN:
		return zapcore.WarnLevel
	case ERROR:
		return zapcore.ErrorLevel
	case FATAL:
		return zapcore.FatalLevel
	case PANIC:
		return zapcore.PanicLevel
	case DPANIC:
		return zapcore.DPanicLevel
	default:
		return zapcore.InfoLevel
	}
}

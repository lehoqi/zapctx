/**
 * @Package zapctx
 * @Time: 2022/2/15 20:32 PM
 * @Author: wuhb
 * @File: rotate.go
 */

package zapctx

import (
	"io"
	"log"
	"path/filepath"

	"github.com/natefinch/lumberjack"
)

type RotateConfig struct {
	MaxAge     int
	MaxBackups int
	Compress   bool
}

func (r *RotateConfig) Init() {
	if r.MaxAge == 0 {
		r.MaxAge = 7
	}
	if r.MaxBackups == 0 {
		r.MaxBackups = 3
	}
}

func Rotate(filePath string, cfg *RotateConfig) io.Writer {
	linkFile, err := filepath.Abs(filePath)
	if err != nil {
		log.Printf("rotateLogger error: %v", err)
		return log.Writer()
	}
	cfg.Init()
	l := &lumberjack.Logger{
		Filename:   linkFile,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  true,
		Compress:   cfg.Compress,
	}
	return l
}

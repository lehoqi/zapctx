/**
 * @Package zapctx
 * @Time: 2022/2/16 20:59 PM
 * @Author: wuhb
 * @File: gorm.go
 */

package zapctx

import (
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func Gorm2(l *Logger) logger.Interface {
	return zapgorm2.New(l.Zap())
}

func Gorm2WithZap(l *zap.Logger) logger.Interface {
	return zapgorm2.New(l)
}

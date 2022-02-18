/**
 * @Package _examples
 * @Time: 2022/2/16 8:21 PM
 * @Author: wuhb
 * @File: main.go
 */

package main

import (
	"context"
	"math/rand"

	"github.com/lehoqi/zapctx"
	"go.uber.org/zap"
)

func main() {
	traceID := "traceID"
	ctx := context.WithValue(context.Background(), traceID, rand.Int())
	logger := zapctx.New(zapctx.WithTraceIDOption(traceID), zapctx.WithDebugOption(true))
	defer logger.Sync()
	zapLog, _ := zap.NewDevelopment()
	defer logger.Sync()
	zapLog = zapLog.Named("zaplog")
	zapLog.Info("zaplog", zapctx.ContextField(traceID, ctx))
	zapLog.Named("sub-zaplog").Info("sub-zaplog")
	zapLog.Info("zaplog")
	logger = logger.Named("main")
	logger.WithContext(ctx).Info("hello")
	logger.WithContext(ctx).Debug("hello")
	logger.WithContext(ctx).Error("hello")
	logger.Named("newlog").Info("hello")
	logger.Debug("hello")
}

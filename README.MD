## Abstract
zapctx is an encapsulation of [zap](https://github.com/uber-go/zap), adding WithContext to the logger.
## Features
- Logging with context, base on zap.
- Integrated with [lumberjack](https://github.com/natefinch/lumberjack), which is a log rotation library.
- Support for [Gorm](https://github.com/go-gorm/gorm), base on [zapgorm2](https://github.com/moul/zapgorm2)

## Installation

```bash
# Go Modules
require github.com/wuhongbing/brickset
```

```shell
go get github.com/wuhongbing/zapctx
```

## Usage

1. Use `zapctx.New()` to create a logger，with `zapctx.WithTraceIDOption` to create a logger with traceID.
```go
package main
import (
	"context"
	"math/rand"

	"github.com/wuhongbing/zapctx"
)
func main() {
	var traceID = "trace_id"
    logger := zapctx.New(zapctx.WithTraceIDOption(traceID),zapctx.WithDebugOption(true),zapctx.WithLogPathOption("/tmp/main.log"))
    defer logger.Sync()
	logger = logger.Named("main")
	ctx := context.WithValue(context.Background(), traceID, rand.Int())
	logger.WithContext(ctx).Info("hello")
	logger.WithContext(ctx).Debug("hello")
	logger.WithContext(ctx).Error("hello")
	logger.Named("sublog").Info("hello")
	logger.Debug("hello")
}
```
```terminal
➜  zapctx cat /tmp/main.log 
2022-02-17 17:40:27.088 INFO    main    _examples/main.go:23    hello   {"trace_id": 5577006791947779410}
2022-02-17 17:40:27.088 DEBUG   main    _examples/main.go:24    hello   {"trace_id": 5577006791947779410}
2022-02-17 17:40:27.088 ERROR   main    _examples/main.go:25    hello   {"trace_id": 5577006791947779410}
2022-02-17 17:40:27.088 INFO    main.sublog     _examples/main.go:26    hello
2022-02-17 17:40:27.088 DEBUG   main    _examples/main.go:27    hello

```
2、Use `zapctx.ContextField` to create a zap field.

```go
package main
import (
    "context"
    "math/rand"
	
    "github.com/wuhongbing/zapctx"
	"go.uber.org/zap"
)

func main() {
	var traceID = "trace_id"
	ctx := context.WithValue(context.Background(), traceID, rand.Int())
	zapLog, _ := zap.NewDevelopment()
	defer zapLog.Sync()
	zapLog = zapLog.Named("zaplog")
	zapLog.Info("zaplog", zapctx.ContextField(traceID, ctx))
}
```

```terminal
➜  zapctx
2022-02-17T17:41:58.173+0800    INFO    zaplog  _examples/main.go:23 zaplog  {"trace_id": 5577006791947779410}

```
3. Use `zapctx.WithCtxDecoderOption` to customize a context decoder.

```go
package main


import (
	"context"
	"math/rand"

	"github.com/wuhongbing/zapctx"
	"go.uber.org/zap"
)

var traceID = "trace_id"

func main() {
	logger := zapctx.New(zapctx.WithCtxDecoderOption(myCtxDecoder(traceID)), zapctx.WithDebugOption(true))
	defer logger.Sync()
	logger = logger.Named("main")
	ctx := context.WithValue(context.Background(), traceID, rand.Int())
	logger.WithContext(ctx).Info("hello")
}
func myCtxDecoder(traceTag string) zapctx.CtxDecoderFunc {
	return func(ctx context.Context, l *zap.Logger) *zap.Logger {
		val := ctx.Value(traceTag)
		if val != nil {
			return l.With(zap.Any(traceTag, val))
		}
		return l
	}
}

```

4. Use `zapctx.WithRotateOption` to back up log file.

```go
package main

import (
	"github.com/wuhongbing/zapctx"
)

func main() {
	logger := zapctx.New(zapctx.WithLogPathOption("/tmp/main.log"), zapctx.WithRotateOption(7,3,true))
	defer logger.Sync()
}
```
5. Use `zapctx.Gorm2` to create a gorm(v2) logger.

```go
package main

import (
	"github.com/wuhongbing/zapctx"
)

func main() {
	logger := zapctx.New()
	defer logger.Sync()
	gormLogger := zapctx.Gorm2(logger.Named("gorm"))
}
```
/**
 * @Package zapctx
 * @Time: 2022/2/16 20:31 PM
 * @Author: wuhb
 * @File: zapctx_test.go
 */

package zapctx

import (
	"context"
	"testing"

	"go.uber.org/zap"
)

var traceID = "traceID"

func TestZapctx(t *testing.T) {
	l := New(WithTraceIDOption(traceID))
	ctx := context.WithValue(context.Background(), traceID, "1-1-1-1-1")
	l.WithContext(ctx).Info("test")
	l.With(zap.String("name", "wuhb")).Info("test2")
	l.Info("test3")

}

func BenchmarkLogger_WithContext(b *testing.B) {
	l := New()
	svr := &tt{logger: l}

	b.Run("withContext", func(b1 *testing.B) {
		ctx1 := context.WithValue(context.Background(), traceID, "1-1-1-1-1")
		b1.ResetTimer()
		b1.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				svr.s1(ctx1)
			}
		})
	})
	b.Run("withoutContext", func(b2 *testing.B) {
		b2.ResetTimer()
		b2.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				svr.s2()
			}
		})
	})
}

type tt struct {
	logger *Logger
}

func (t *tt) s1(ctx context.Context) {
	t.logger.WithContext(ctx).Info("test")
}
func (t *tt) s2() {
	t.logger.With(zap.String(traceID, "1-1-1-1-1")).Info("test")
}

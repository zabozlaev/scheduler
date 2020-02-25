
package http

import (
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func (a *adapter) loggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		t1 := time.Now()
		defer func() {
			a.logger.Info("Served",
				zap.String("proto", r.Proto),
				zap.String("path", r.URL.Path),
				zap.Duration("lat", time.Since(t1)),
				zap.Int("status", ww.Status()),
				zap.Int("size", ww.BytesWritten()),
			)
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
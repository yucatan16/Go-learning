package middlewares

import (
	"context"
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

type traceIDkey struct{}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDkey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDkey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

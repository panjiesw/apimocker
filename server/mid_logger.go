package server

import (
	"fmt"
	"net/http"
	"time"

	"runtime/debug"

	"github.com/mgutz/logxi/v1"
)

type RequestLogger struct {
	log.Logger
	id   string
	args []interface{}
}

// Trace logs a debug entry.
func (l *RequestLogger) Trace(msg string, args ...interface{}) {
	l.Logger.Trace(msg, append([]interface{}{"req_id", l.id}, args...)...)
}

// Debug logs a debug entry.
func (l *RequestLogger) Debug(msg string, args ...interface{}) {
	l.Logger.Debug(msg, append([]interface{}{"req_id", l.id}, args...)...)
}

// Info logs an info entry.
func (l *RequestLogger) Info(msg string, args ...interface{}) {
	l.Logger.Info(msg, append([]interface{}{"req_id", l.id}, args...)...)
}

// Warn logs a warn entry.
func (l *RequestLogger) Warn(msg string, args ...interface{}) error {
	return l.Logger.Warn(msg, append([]interface{}{"req_id", l.id}, args...)...)
}

// Error logs an error entry.
func (l *RequestLogger) Error(msg string, args ...interface{}) error {
	return l.Logger.Error(msg, append([]interface{}{"req_id", l.id}, args...)...)
}

// Fatal logs a fatal entry then panics.
func (l *RequestLogger) Fatal(msg string, args ...interface{}) {
	l.Logger.Fatal(msg, append([]interface{}{"req_id", l.id}, args...)...)
}

// Log logs a leveled entry.
func (l *RequestLogger) Log(level int, msg string, args []interface{}) {
	l.Logger.Log(level, msg, append([]interface{}{"req_id", l.id}, args...))
}

func (l *RequestLogger) Start(r *http.Request) {

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	l.args = append(l.args, []interface{}{
		"http_scheme", scheme,
		"http_proto", r.Proto,
		"http_method", r.Method,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
		"uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
	}...)
	l.Info("started", l.args...)
}

func (l *RequestLogger) End(status, bytes int, elapsed time.Duration) {
	l.args = append(l.args, []interface{}{
		"resp_status", status,
		"resp_bytes_length", bytes,
		"resp_elasped_ms", float64(elapsed.Nanoseconds()) / 1000000.0,
	}...)
	l.Info("end", l.args...)
}

func NewRequestLogger(id string) *RequestLogger {
	logger := log.New("request")
	return &RequestLogger{Logger: logger, id: id, args: make([]interface{}, 0)}
}

// Ported from lg's middleware, source:
// https://github.com/pressly/lg/blob/master/middleware.go

func (s *Server) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := r.Context().Value(RootCtxKey).(*RootCtx)
		logger := NewRequestLogger(rctx.ID)
		logger.Start(r)
		ww := NewWrapResponseWriter(w, r.ProtoMajor)

		t1 := time.Now()
		defer func() {
			t2 := time.Now()

			// Recover and record stack traces in case of a panic
			if rec := recover(); rec != nil {
				logger.Error("that happens", "panic", fmt.Sprintf("%+v", rec), "stack", string(debug.Stack()))
				http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

			// Log the entry, the request is complete.
			logger.End(ww.Status(), ww.BytesWritten(), t2.Sub(t1))
		}()

		rctx.L = logger
		next.ServeHTTP(ww, r)
	})
}

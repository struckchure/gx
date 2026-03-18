package internals

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

type plainHandler struct{}

func (h *plainHandler) Enabled(_ context.Context, _ slog.Level) bool { return true }
func (h *plainHandler) WithAttrs(_ []slog.Attr) slog.Handler         { return h }
func (h *plainHandler) WithGroup(_ string) slog.Handler              { return h }
func (h *plainHandler) Handle(_ context.Context, r slog.Record) error {
	fmt.Fprintf(os.Stdout, "[%s] [%s] %s\n",
		r.Time.Format(time.DateTime),
		r.Level,
		r.Message,
	)
	return nil
}

var (
	Logger     *slog.Logger
	loggerOnce sync.Once
)

func init() {
	loggerOnce.Do(func() {
		Logger = slog.New(&plainHandler{})
		// slog.SetDefault(Logger)
	})
}

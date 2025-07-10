// Il pacchetto log provvede un'interfaccia di logging.
// I livelli di log sono:
//   - Debug (Solo nelle build di sviluppo)
//   - Info
//   - Warn
//   - Error
//   - Panic (Chiama panic() dopo aver loggato)
//   - Fatal (Chiama os.Exit(1) dopo aver loggato)
package log

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/MrRainbow0704/ProgettoIride/internal/version"
)

var (
	logger    *slog.Logger
	WebLogger *log.Logger
)

const (
	LevelPanic = slog.Level(10)
	LevelFatal = slog.Level(12)
)

func init() {
	var handler slog.Handler = slog.NewJSONHandler(
		os.Stderr,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		},
	)
	if version.IsDev() {
		handler = slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		)
	}
	logger = slog.New(handler)
	WebLogger = slog.NewLogLogger(logger.WithGroup("web").Handler(), slog.LevelInfo)
}

func Debug(msg string, args map[string]any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelDebug, msg, vals...)
}

func Debugf(msg string, args map[string]any, a ...any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelDebug, fmt.Sprintf(msg, a), vals...)
}

func Info(msg string, args map[string]any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelInfo, msg, vals...)
}

func Infof(msg string, args map[string]any, a ...any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelInfo, fmt.Sprintf(msg, a), vals...)
}

func Warn(msg string, args map[string]any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelWarn, msg, vals...)
}

func Warnf(msg string, args map[string]any, a ...any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelWarn, fmt.Sprintf(msg, a), vals...)
}

func Error(msg string, args map[string]any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelError, msg, vals...)
}

func Errorf(msg string, args map[string]any, a ...any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), slog.LevelError, fmt.Sprintf(msg, a), vals...)
}

func Panic(msg string, args map[string]any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), LevelPanic, msg, vals...)
	panic(msg)
}

func Panicf(msg string, args map[string]any, a ...any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), LevelPanic, fmt.Sprintf(msg, a), vals...)
	panic(msg)
}

func Fatal(msg string, args map[string]any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), LevelFatal, msg, vals...)
	os.Exit(1)
}

func Fatalf(msg string, args map[string]any, a ...any) {
	vals := []slog.Attr{}
	for k, v := range args {
		vals = append(vals, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	logger.LogAttrs(context.Background(), LevelFatal, fmt.Sprintf(msg, a), vals...)
	os.Exit(1)
}

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
	"io"
	"log/slog"
	"os"
	"runtime"

	"github.com/MrRainbow0704/ProgettoIride/internal/version"
)

var logger *slog.Logger

const (
	LevelPanic = slog.Level(10)
	LevelFatal = slog.Level(12)
)

func LoadLogger(writers ...io.Writer) {
	logger = slog.New(
		slog.NewJSONHandler(
			io.MultiWriter(writers...),
			&slog.HandlerOptions{
				ReplaceAttr: replaceAttr,
				Level:       slog.LevelInfo,
			},
		),
	)
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		switch a.Value.Any().(slog.Level) {
		case LevelPanic:
			a.Value = slog.StringValue("PANIC")
		case LevelFatal:
			a.Value = slog.StringValue("FATAL")
		}
	}
	return a
}

func getSource() slog.Attr {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return slog.Attr{}
	}

	funcName := "unknown"
	if details := runtime.FuncForPC(pc); details != nil {
		funcName = details.Name()
	}

	return slog.GroupAttrs("source",
		slog.String("function", funcName),
		slog.String("file", file),
		slog.Int("line", line),
	)
}

func Debug(msg string) {
	if !version.IsDev() {
		return
	}
	logger.LogAttrs(
		context.Background(), slog.LevelDebug,
		msg, getSource(),
	)
}

func Debugf(msg string, a ...any) {
	if !version.IsDev() {
		return
	}
	logger.LogAttrs(
		context.Background(), slog.LevelDebug,
		fmt.Sprintf(msg, a...), getSource(),
	)
}

func Info(msg string) {
	logger.LogAttrs(
		context.Background(), slog.LevelInfo,
		msg, getSource(),
	)
}

func Infof(msg string, a ...any) {
	logger.LogAttrs(
		context.Background(), slog.LevelInfo,
		fmt.Sprintf(msg, a...), getSource(),
	)
}

func Warn(msg string) {
	logger.LogAttrs(
		context.Background(), slog.LevelWarn,
		msg, getSource(),
	)
}

func Warnf(msg string, a ...any) {
	logger.LogAttrs(
		context.Background(), slog.LevelWarn,
		fmt.Sprintf(msg, a...), getSource(),
	)
}

func Error(msg string) {
	logger.LogAttrs(
		context.Background(), slog.LevelError,
		msg, getSource(),
	)
}

func Errorf(msg string, a ...any) {
	logger.LogAttrs(
		context.Background(), slog.LevelError,
		fmt.Sprintf(msg, a...), getSource(),
	)
}

func Panic(msg string) {
	logger.LogAttrs(
		context.Background(), LevelPanic,
		msg, getSource(),
	)
	panic(msg)
}

func Panicf(msg string, a ...any) {
	logger.LogAttrs(
		context.Background(), LevelPanic,
		fmt.Sprintf(msg, a...), getSource(),
	)
	panic(msg)
}

func Fatal(msg string) {
	logger.LogAttrs(
		context.Background(), LevelFatal,
		msg, getSource(),
	)
	os.Exit(1)
}

func Fatalf(msg string, a ...any) {
	logger.LogAttrs(
		context.Background(), LevelFatal,
		fmt.Sprintf(msg, a...), getSource(),
	)
	os.Exit(1)
}

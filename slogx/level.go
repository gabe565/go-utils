package slogx

import (
	"bytes"
	"context"
	"log/slog"
	"strconv"
	"strings"
)

type Level slog.Level

const (
	LevelTrace = Level(slog.LevelDebug) - 1
	LevelDebug = Level(slog.LevelDebug)
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)
)

func LevelStrings() []string {
	return []string{
		LevelTrace.String(),
		LevelDebug.String(),
		LevelInfo.String(),
		LevelWarn.String(),
		LevelError.String(),
	}
}

func (l *Level) UnmarshalText(text []byte) error {
	if bytes.HasPrefix(bytes.ToLower(text), []byte("trace")) {
		*l = LevelTrace
		if i := bytes.IndexAny(text, "+-"); i >= 0 {
			offset, err := strconv.Atoi(string(text[i:]))
			if err != nil {
				return err
			}
			*l += Level(offset)
		}
		return nil
	}

	var slogLevel slog.Level
	if err := slogLevel.UnmarshalText(text); err != nil {
		return err
	}
	*l = Level(slogLevel)
	return nil
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l Level) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, l.String()), nil
}

func (l *Level) UnmarshalJSON(i []byte) error {
	s, err := strconv.Unquote(string(i))
	if err != nil {
		return err
	}
	return l.Set(s)
}

func (l *Level) Set(s string) error {
	return l.UnmarshalText([]byte(s))
}

func (l Level) Type() string {
	return "string"
}

func (l Level) String() string {
	if l == LevelTrace {
		return "trace"
	}
	return strings.ToLower(l.Level().String())
}

func (l Level) Level() slog.Level {
	return slog.Level(l)
}

func Trace(msg string, args ...any) {
	LoggerTrace(slog.Default(), msg, args...)
}

func LoggerTrace(logger *slog.Logger, msg string, args ...any) {
	logger.Log(context.Background(), LevelTrace.Level(), msg, args...)
}

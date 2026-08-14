// Package logger provides a minimal leveled logger used by fixture packages.
package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

type Logger struct {
	out   io.Writer
	level Level
}

func New() *Logger {
	return &Logger{out: os.Stderr, level: LevelInfo}
}

func (l *Logger) SetLevel(lvl Level) { l.level = lvl }

func (l *Logger) log(lvl Level, msg string) {
	if lvl < l.level {
		return
	}
	fmt.Fprintf(l.out, "%s [%d] %s\n", time.Now().Format(time.RFC3339), lvl, msg)
}

func (l *Logger) Info(msg string)  { l.log(LevelInfo, msg) }
func (l *Logger) Warn(msg string)  { l.log(LevelWarn, msg) }
func (l *Logger) Error(msg string) { l.log(LevelError, msg) }
func (l *Logger) Debug(msg string) { l.log(LevelDebug, msg) }

// Infof formats and logs at INFO level.
func (l *Logger) Infof(format string, args ...any) {
	l.log(LevelInfo, fmt.Sprintf(format, args...))
}

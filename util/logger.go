package util

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lj "gopkg.in/natefinch/lumberjack.v2"
)

// Logger is a wrapper of uber/zap logger with dynamic log level.
type Logger struct {
	logL     *zap.Logger
	logS     *zap.SugaredLogger
	minLevel *zap.AtomicLevel
}

var (
	getEncoderConfig = func(lvlEnc zapcore.LevelEncoder) *zapcore.EncoderConfig {
		return &zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    lvlEnc,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	}
	// check if the level is greater than or equal to error
	allLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	// log encoders
	consoleEncoder = zapcore.NewConsoleEncoder(*getEncoderConfig(zapcore.CapitalColorLevelEncoder))
	jsonEncoder    = zapcore.NewJSONEncoder(*getEncoderConfig(zapcore.LowercaseLevelEncoder))
	// std log writers
	writeStdout = zapcore.AddSync(os.Stdout)
	writeStderr = zapcore.AddSync(os.Stderr)
)

// NewLogger returns a Logger with given log path and debug mode.
func NewLogger(fileName string, debug bool) *Logger {
	// log enablers
	minLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	normalLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= minLevel.Level()
	})
	errorLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && lvl >= minLevel.Level()
	})

	// combine core for logger
	cores := []zapcore.Core{
		zapcore.NewCore(consoleEncoder, writeStdout, normalLevelEnabler),
		zapcore.NewCore(consoleEncoder, writeStderr, errorLevelEnabler),
	}
	// set log file path as per parameters
	if logPath := strings.TrimSpace(fileName); len(logPath) > 0 {
		logFile := zapcore.AddSync(&lj.Logger{
			Filename: logPath,
			MaxSize:  10, // megabytes
		})
		cores = append(cores, zapcore.NewCore(jsonEncoder, logFile, allLevelEnabler))
	}

	// combine option for logger
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	}
	// set debug mode as per parameters
	if debug {
		options = append(options, zap.Development())
	}

	// build the logger
	logL := zap.New(zapcore.NewTee(cores...), options...)
	return &Logger{
		logL:     logL,
		logS:     logL.Sugar(),
		minLevel: &minLevel,
	}
}

// Logger returns a zap logger inside the wrapper.
func (l *Logger) Logger() *zap.Logger {
	return l.logL
}

// LoggerSugared returns a sugared zap logger inside the wrapper.
func (l *Logger) LoggerSugared() *zap.SugaredLogger {
	return l.logS
}

// SetLogLevel sets the log level of loggers inside the wrapper.
func (l *Logger) SetLogLevel(level string) {
	if err := l.minLevel.UnmarshalText([]byte(level)); err != nil {
		l.logL.DPanic("fail to set log level", zap.Error(err))
	}
}

// GetLogLevel returns the log level of loggers inside the wrapper.
func (l *Logger) GetLogLevel() zapcore.Level {
	return l.minLevel.Level()
}

package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Converts a logrus.Logger to a zap.Logger
func LogrusToZap(logger *logrus.Logger) *zap.Logger {
	// Create encoder config based on logrus formatter settings
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Check if we're using JSON formatter
	var encoder zapcore.Encoder
	if _, ok := logger.Formatter.(*logrus.JSONFormatter); ok {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Set log levels
	zapLevel := logrusLevelToZapLevel(logger.Level)

	// Use same output as logrus
	writeSyncer := zapcore.AddSync(logger.Out)

	// Create the core
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel)

	// Build the logger
	return zap.New(core)
}

// Converts a logrus.Entry to a zap.Logger
func LogrusEntryToZap(entry *logrus.Entry) *zap.Logger {
	zapLogger := LogrusToZap(entry.Logger)

	// Add fields from logrus.Entry
	fields := make([]zap.Field, 0, len(entry.Data))
	for k, v := range entry.Data {
		fields = append(fields, zap.Any(k, v))
	}

	return zapLogger.With(fields...)
}

// Convert logrus level to zap level
func logrusLevelToZapLevel(level logrus.Level) zap.AtomicLevel {
	switch level {
	case logrus.TraceLevel, logrus.DebugLevel:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case logrus.InfoLevel:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case logrus.WarnLevel:
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case logrus.ErrorLevel:
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case logrus.FatalLevel:
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case logrus.PanicLevel:
		return zap.NewAtomicLevelAt(zapcore.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

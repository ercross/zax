package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"io"
)

// Level represents the severity of the log message.
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LoggerConfig struct {
	Level       Level
	Destination io.WriteCloser
	ServiceName string
}

type Logger struct {
	zap    *zap.Logger
	config LoggerConfig
}

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value interface{}
}

func NewLogger(config LoggerConfig) (*Logger, error) {

	w := zapcore.AddSync(config.Destination)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zapcore.Level(config.Level),
	)

	zapLogger := zap.New(core)
	return &Logger{zap: zapLogger, config: config}, nil
}

// NewSilentLogger creates a logger that silences all log output
func NewSilentLogger() (*Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.FatalLevel + 1)
	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return &Logger{zap: zapLogger, config: LoggerConfig{}}, nil
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, toZapFields(fields)...)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, toZapFields(fields)...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, toZapFields(fields)...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, toZapFields(fields)...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, toZapFields(fields)...)
}

func (l *Logger) Flush() error {
	return l.zap.Sync()
}

// toZapFields converts our custom Field type to zap.Field, keeping zap.Field encapsulated.
func toZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger provides an interface for logging in CoC
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Sync() error
}

type logger struct {
	sugar *zap.SugaredLogger
}

func NewLogger() (Logger, error) {
	//zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config := zap.NewProductionConfig()
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)
	//config.OutputPaths = []string{"stdout", "./logs/" + logFile}
	z, err := config.Build(zap.AddCaller())
	if err != nil {
		return nil, err
	}
	return &logger{sugar: z.Sugar()}, err
}

func New() Logger {
	z, _ := zap.NewProduction()
	l := &logger{sugar: z.Sugar()}
	return l
}

func NewWithConfig(cfg *zap.Config) Logger {
	z, _ := cfg.Build()
	l := &logger{sugar: z.Sugar()}
	return l
}

// SetLevel sets the log level for logger
func (l *logger) SetLevel(level int8) error {
	config := zap.NewProductionConfig()
	a := zap.NewAtomicLevel()
	a.SetLevel(zapcore.Level(level))
	z, err := config.Build(zap.AddCaller())
	if err != nil {
		return err
	}
	l.sugar = z.Sugar()
	return nil
}

func (l *logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

func (l *logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *logger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

func (l *logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *logger) Sync() error {
	return l.sugar.Sync()
}

package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type loggerLogrus struct {
	log *logrus.Logger
}

func NewLogger() Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		ForceColors:     true,
	})

	log.SetLevel(logrus.InfoLevel)
	log.SetOutput(os.Stdout)

	logrus := loggerLogrus{
		log: log,
	}

	return &logrus
}

func (l *loggerLogrus) Info(args ...any) {
	l.log.Info(args...)
}

func (l *loggerLogrus) Infof(format string, args ...any) {
	l.log.Infof(format, args...)
}

func (l *loggerLogrus) Warn(args ...any) {
	l.log.Warn(args...)
}

func (l *loggerLogrus) Warnf(format string, args ...any) {
	l.log.Warnf(format, args...)
}

func (l *loggerLogrus) Error(args ...any) {
	l.log.Error(args...)
}

func (l *loggerLogrus) Errorf(format string, args ...any) {
	l.log.Errorf(format, args...)
}

func (l *loggerLogrus) Fatal(args ...any) {
	l.log.Fatal(args...)
}

func (l *loggerLogrus) Fatalf(format string, args ...any) {
	l.log.Fatalf(format, args...)
}

func (l *loggerLogrus) Debug(args ...any) {
	l.log.Debug(args...)
}

func (l *loggerLogrus) Debugf(format string, args ...any) {
	l.log.Debugf(format, args...)
}

func (l *loggerLogrus) WithField(key string, value any) Logger {
	return &LoggerEntry{
		entry: l.log.WithField(key, value),
	}
}

func (l *loggerLogrus) WithFields(args map[string]any) Logger {
	return &LoggerEntry{
		entry: l.log.WithFields(args),
	}
}

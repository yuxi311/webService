package logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/pkg/utils"
)

type Options struct {
	Mode       string
	Level      string
	Path       string
	Format     string
	MaxSize    int
	MaxBackups int
}

type MillisecondsFormatter struct {
	logrus.Formatter
}

func Init() error {
	cfg := config.Log()
	opts := Options{
		Mode:       cfg.Mode,
		Level:      cfg.Level,
		Path:       utils.ToFilePath(cfg.File),
		Format:     cfg.Format,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
	}

	level, err := logrus.ParseLevel(opts.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	if opts.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&MillisecondsFormatter{
			&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339Nano,
				DisableColors:   true,
			},
		})
	}

	if opts.Mode == "file" && len(opts.Path) > 0 {
		log := &lumberjack.Logger{
			Filename:   opts.Path,
			MaxSize:    opts.MaxSize,
			MaxBackups: opts.MaxBackups,
		}
		logrus.SetOutput(log)
	} else {
		logrus.SetOutput(os.Stdout)
	}
	return nil
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

package database

import (
	"context"
	"errors"
	"time"

	"github.com/delta/arcadia-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var databaseLogger *logrus.Logger

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

func newDBLogger() *logger {
	return &logger{
		SkipErrRecordNotFound: true,
		Debug:                 true,
	}
}

func initDatabaseLogger(filename string) {
	databaseLogger = utils.NewLogger(filename)
}

func (l *logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	databaseLogger.WithContext(ctx).Infof(s, args...)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	databaseLogger.WithContext(ctx).Warnf(s, args...)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	databaseLogger.WithContext(ctx).Errorf(s, args...)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		databaseLogger.WithContext(ctx).WithFields(logrus.Fields{
			"sql":     sql,
			"elapsed": elapsed,
		}).Log(logrus.ErrorLevel, err)
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		databaseLogger.WithContext(ctx).WithFields(logrus.Fields{
			"sql":     sql,
			"elapsed": elapsed,
		}).Log(logrus.WarnLevel)
		return
	}

	if l.Debug {
		databaseLogger.WithContext(ctx).WithFields(logrus.Fields{
			"sql":     sql,
			"elapsed": elapsed,
		}).Log(logrus.InfoLevel)
	}
}

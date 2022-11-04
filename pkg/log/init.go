package log

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	logger = log.New()

	debugOn = os.Getenv("LOG_DEBUG")
)

const logFormat = "[%v][%v][%v][ctx=%v]: %v\n"

type logFormatter struct{}

func (lf *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf(
		logFormat,
		entry.Time.Format(time.RFC3339),
		strings.ToUpper(entry.Level.String()),
		entry.Data["file"],
		entry.Context,
		entry.Message,
	)), nil
}

func init() {
	logger.SetFormatter(new(logFormatter))
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)

	if debugOn == "true" {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetLevel(log.InfoLevel)
	}
}

var entry = logger.WithFields(log.Fields{})

func Debug(msg string, args ...interface{}) {
	entry.Data["file"] = fileInfo(2)
	entry.Debugf(msg, args...)
}

func Info(msg string, args ...interface{}) {
	entry.Data["file"] = fileInfo(2)
	entry.Infof(msg, args...)
}

func InfoCtx(ctx context.Context, msg string, args ...interface{}) {
	entry.Context = ctx
	Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	entry.Data["file"] = fileInfo(2)
	entry.Warnf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	entry.Data["file"] = fileInfo(2)
	entry.Errorf(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	entry.Data["file"] = fileInfo(2)
	entry.Fatalf(msg, args...)
}

const serviceName = "mlib"

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<nil>"
		line = 1
	} else {
		file = file[strings.Index(file, serviceName):]
		slash := strings.IndexAny(file, "/")
		file = file[slash+1:]
	}
	return fmt.Sprintf("%s:%d", file, line)
}

package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

const LogLevel = logrus.DebugLevel

// GetLogger Init initialize logger.
func GetLogger(logPath string, level int, printLogsToStdOut bool) (Logger, error) {
	log := logrus.New()

	// l.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	log.SetLevel(LogLevel)

	logFile, err := os.Create(logPath)
	if err != nil {
		return Logger{}, fmt.Errorf("wailed to init logger error: %w", err)
	}

	mw := io.MultiWriter(logFile)
	if printLogsToStdOut {
		mw = io.MultiWriter(os.Stdout, logFile)
	}

	log.SetOutput(mw)
	log.SetLevel(logrus.Level(level))

	return Logger{log}, nil
}

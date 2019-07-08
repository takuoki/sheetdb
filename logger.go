package sheetdb

import (
	"log"
	"os"

	"github.com/takuoki/slack-alert/slacka"
)

// Logger is an interface for logging of this package.
type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

// LogLevel represents the level of logging when using the default logger.
type LogLevel int

// Levels of logging.
const (
	LevelDebug LogLevel = iota + 1
	LevelInfo
	LevelError
	LevelNoLogging
)

var (
	logger Logger = &defaultLogger{
		out: log.New(os.Stdout, "", log.LstdFlags),
		err: log.New(os.Stderr, "", log.LstdFlags),
	}
	logLevel = LevelInfo
)

// SetLogger sets a logger for this package.
// By default, logs are output only to standard output and standard error output.
func SetLogger(l Logger) {
	logger = l
}

// SetLogLevel sets a log level for this package.
// If you don't use the default logger or SlackLogger, log level are meaningless.
// The default log level is "INFO".
func SetLogLevel(l LogLevel) {
	logLevel = l
}

type defaultLogger struct {
	out *log.Logger
	err *log.Logger
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	if logLevel > LevelDebug {
		return
	}
	l.out.Printf(format, v...)
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	if logLevel > LevelInfo {
		return
	}
	l.out.Printf(format, v...)
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	if logLevel > LevelError {
		return
	}
	l.err.Printf(format, v...)
}

// SlackLogger is a logget to output log to slack.
// Logs are also output to standard output and standard error output at the same time,
// and you can control each log level.
type SlackLogger struct {
	logger        *defaultLogger
	slackaClient  *slacka.Client
	slackLogLevel LogLevel
}

// NewSlackLogger returns a new SlackLogger.
func NewSlackLogger(projectName, serviceName, iconEmoji, errorURL string,
	slackLogLevel LogLevel) *SlackLogger {
	client := slacka.New(projectName, serviceName, iconEmoji)
	client.SetErrorURL(errorURL)
	return &SlackLogger{
		logger: &defaultLogger{
			out: log.New(os.Stdout, "", log.LstdFlags),
			err: log.New(os.Stderr, "", log.LstdFlags),
		},
		slackaClient:  client,
		slackLogLevel: slackLogLevel,
	}
}

// Debugf outputs a log whose log level is Debug or higher.
func (l *SlackLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
	if l.slackLogLevel > LevelDebug {
		return
	}
	if err := l.slackaClient.Debugf(format, v...); err != nil {
		l.logger.Errorf("Unable to send alert to slack: %v", err)
	}
}

// Infof outputs a log whose log level is Info or higher.
func (l *SlackLogger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
	if l.slackLogLevel > LevelInfo {
		return
	}
	if err := l.slackaClient.Infof(format, v...); err != nil {
		l.logger.Errorf("Unable to send alert to slack: %v", err)
	}
}

// Errorf outputs a log whose log level is Error or higher.
func (l *SlackLogger) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
	if l.slackLogLevel > LevelError {
		return
	}
	if err := l.slackaClient.Errorf(format, v...); err != nil {
		l.logger.Errorf("Unable to send alert to slack: %v", err)
	}
}

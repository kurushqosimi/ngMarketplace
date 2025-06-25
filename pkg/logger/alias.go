package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

// Level wrapper so that the custom logger level would be the same as a zerolog logger level
type (
	Level = zerolog.Level
)

// CustomLogger wrapper of zerolog.Logger
type CustomLogger struct {
	logger zerolog.Logger
}

const (
	TraceLevel = zerolog.TraceLevel
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	FatalLevel = zerolog.FatalLevel
	PanicLevel = zerolog.PanicLevel
	NoLevel    = zerolog.NoLevel
)

// We need this variables/functions, so we can configure global level of logging.
var (
	SetGlobalLevel = zerolog.SetGlobalLevel

	// ConsoleWriter comfortable writer for beautiful console output.
	ConsoleWriter = zerolog.ConsoleWriter{Out: os.Stderr}
)

// NewTextHandler returns "text" (ConsoleWriter) logger.
func NewTextHandler(w io.Writer) zerolog.Logger {
	return zerolog.New(w).With().Timestamp().Logger()
}

// NewJSONHandler returns "JSON" logger.
func NewJSONHandler(w io.Writer) zerolog.Logger {
	if w == nil {
		w = os.Stderr
	}
	return zerolog.New(w).With().Timestamp().Logger()

}

// NewFileWriter configures `lumberjack.Logger` for file rotation.
func NewFileWriter(config *LoggerOptions) io.Writer {
	return &lumberjack.Logger{
		Filename:   config.LogFilePath,
		MaxSize:    config.LogFileMaxSizeMB,
		MaxAge:     config.LogFileMaxAgeDays,
		MaxBackups: config.LogFileMaxBackups,
		Compress:   config.LogFileCompress,
	}
}

// Debug implements Debug log level with zerolog.
func (l *CustomLogger) Debug(message interface{}, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Debug().Msg(fmt.Sprintf("%v", message))
	} else {
		l.logger.Debug().Msgf(fmt.Sprintf("%v", message), args...)
	}
}

// Info implements Info log level with zerolog.
func (l *CustomLogger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

// Warn implements Warn log level with zerolog.
func (l *CustomLogger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Warn().Msg(message)
	} else {
		l.logger.Warn().Msgf(message, args...)
	}
}

// Error implements Error log level with zerolog.
func (l *CustomLogger) Error(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Error().Msg(message)
	} else {
		l.logger.Error().Msgf(message, args...)
	}
}

// Fatal implements Fatal log level with zerolog.
func (l *CustomLogger) Fatal(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Fatal().Msg(message)
	} else {
		l.logger.Fatal().Msgf(message, args...)
	}
}

// Printf implements Printf log level with zerolog.
func (l *CustomLogger) Printf(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Print(message)
	} else {
		l.logger.Printf(message, args...)
	}
}

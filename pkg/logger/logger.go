package logger

import (
	"io"
	"sync"
)

// Logger — interface that wraps the concrete implementation.
type Logger interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
	Printf(message string, args ...interface{})
}

var (
	once     sync.Once
	instance *CustomLogger
)

const (
	defaultLevel             = InfoLevel
	defaultAddSource         = true
	defaultIsJSON            = true
	defaultSetDefault        = true
	defaultLogFile           = ""
	defaultLogFileMaxSizeMB  = 10
	defaultLogFileMaxBackups = 3
	defaultLogFileMaxAgeDays = 14
	defaultLogFileCompress   = false
)

// New collects all option in LoggerOptions, creates single CustomLogger (Singleton)
func New(opts ...LoggerOption) *CustomLogger {
	once.Do(func() {
		config := &LoggerOptions{
			Level:             defaultLevel,
			AddSource:         defaultAddSource,
			IsJSON:            defaultIsJSON,
			SetDefault:        defaultSetDefault,
			LogFilePath:       defaultLogFile,
			LogFileMaxSizeMB:  defaultLogFileMaxSizeMB,
			LogFileMaxBackups: defaultLogFileMaxBackups,
			LogFileMaxAgeDays: defaultLogFileMaxAgeDays,
			LogFileCompress:   defaultLogFileCompress,
		}

		for _, opt := range opts {
			opt(config)
		}

		var logOutput io.Writer
		if config.LogFilePath != "" {
			logOutput = NewFileWriter(config)
		}

		var baseLogger = NewJSONHandler(logOutput)
		if !config.IsJSON {
			baseLogger = NewTextHandler(logOutput)
		}

		if config.SetDefault {
			SetGlobalLevel(config.Level)
		}

		if config.AddSource {
			baseLogger = baseLogger.With().CallerWithSkipFrameCount(3).Logger()
		}

		// Инициализируем CustomLogger
		instance = &CustomLogger{
			logger: baseLogger,
		}
	})

	return instance
}

// LoggerOption — function for modification LoggerOptions
type LoggerOption func(options *LoggerOptions)

// WithLevel - function that defines the log level
func WithLevel(level string) LoggerOption {
	return func(o *LoggerOptions) {
		var l Level
		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = InfoLevel
		}

		o.Level = l
	}
}

// WithAddSource - function that adds the source of the log
func WithAddSource(addSource bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.AddSource = addSource
	}
}

// WithIsJSON - function that defines that logs would be in JSON format
func WithIsJSON(isJSON bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.IsJSON = isJSON
	}
}

// WithSetDefault - sets logger default
func WithSetDefault(setDefault bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.SetDefault = setDefault
	}
}

// WithLogFilePath - sets where to save log file
func WithLogFilePath(logFilePath string) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFilePath = logFilePath
	}
}

// WithLogFileMaxSizeMB - sets the maximum file size of a log file
func WithLogFileMaxSizeMB(maxSize int) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileMaxSizeMB = maxSize
	}
}

// WithLogFileMaxBackups - sets the number of maximum backups
func WithLogFileMaxBackups(maxBackups int) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileMaxBackups = maxBackups
	}
}

// WithLogFileCompress - sets the compression
func WithLogFileCompress(compression bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.LogFileCompress = compression
	}
}

// LoggerOptions — набор опций
type LoggerOptions struct {
	Level             Level
	AddSource         bool
	IsJSON            bool
	SetDefault        bool
	LogFilePath       string
	LogFileMaxSizeMB  int
	LogFileMaxBackups int
	LogFileMaxAgeDays int
	LogFileCompress   bool
}

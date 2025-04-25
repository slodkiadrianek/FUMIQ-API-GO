package utils

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Logger struct {
	Logger zerolog.Logger
}

func (l Logger) CreateLogger() Logger {
	logDir := "./logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	actualDate := time.Now()
	fileName := actualDate.Format("2006-01-02")
	file, err := os.OpenFile(logDir+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	multi := zerolog.MultiLevelWriter(os.Stderr, file)

	logger := zerolog.New(multi).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	logger.Warn().Msg("Logger created")
	return Logger{Logger: logger}
}
func (l Logger) Info(message string, data ...interface{}) {
	l.Logger.Info().Msg(message)
}
func (l Logger) Debug(message string, data ...interface{}) {
	l.Logger.Debug().Msg(message)
}
func (l Logger) Error(message string, data ...interface{}) {
	l.Logger.Error().Msg(message)
}

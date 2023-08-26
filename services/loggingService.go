package services

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Fields(fields map[string]interface{})
	Error(err error)
	Debug(msg string)
	Info(msg string)
}

type consoleLogger struct {
	logger zerolog.Logger
}

func NewConsoleLogger(service string) Logger {
	l := new(consoleLogger)
	l.logger = zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		Str("service", service).
		Logger()
	return l
}

func (cl *consoleLogger) Fields(fields map[string]interface{}) {
	cl.logger.Info().Fields(fields).Send()
}

func (cl *consoleLogger) Error(err error) {
	cl.logger.Err(err)
}

func (cl *consoleLogger) Debug(msg string) {
	cl.logger.Debug().Msg(msg)
}

func (cl *consoleLogger) Info(msg string) {
	cl.logger.Info().Msg(msg)
}

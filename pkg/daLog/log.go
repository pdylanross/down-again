package daLog

import (
	"time"

	"go.uber.org/zap"
)

// Logger interface for DA
type Logger interface {
	LogSuccess(pingTarget string, duration time.Duration)
	LogFail(pingTarget string)
}

// LoggerOptions for log construction
type LoggerOptions struct {
}

type loggerImpl struct {
	log *zap.Logger
}

// CreateLogger creates a DA logger
func CreateLogger(options *LoggerOptions) (Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.DisableCaller = true

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &loggerImpl{log: log}, nil
}

func (log loggerImpl) LogSuccess(pingTarget string, duration time.Duration) {
	log.log.Info(
		pingTarget,
		zap.Bool("s", true),
		zap.Duration("d", duration),
	)
}

func (log loggerImpl) LogFail(pingTarget string) {
	log.log.Info(
		pingTarget,
		zap.Bool("s", false),
		zap.Duration("d", 0),
	)
}

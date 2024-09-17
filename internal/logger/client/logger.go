package client

import "go.uber.org/zap"

type ClientLogger struct {
	*zap.Logger
}

func Initialize(level string) (*ClientLogger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &ClientLogger{logger}, nil
}

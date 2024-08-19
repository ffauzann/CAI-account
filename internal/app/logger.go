package app

import (
	"log"

	"github.com/ffauzann/CAI-account/internal/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Format string
	Zap    *zap.Logger
}

func (l *Logger) init() (err error) {
	config := zap.NewDevelopmentConfig()

	config.DisableStacktrace = false
	config.Encoding = l.Format
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		CallerKey:     "caller",
		NameKey:       zapcore.OmitKey,
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: zapcore.OmitKey,

		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// TODO: Store config.
	// TODO: Pass config to interceptor.
	// TODO: Define zincsearch interceptor.
	// BLOCKER: current util.LogContext uses different logger with new core later.

	l.Zap, err = config.Build()
	if err != nil {
		log.Fatalln(err)
		return
	}

	zap.ReplaceGlobals(l.Zap)
	util.SetLogger(l.Zap)
	return
}

// go-ecommerce-backend-api/pkg/logger/logger.go
package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(cfg *setting.LoggerSetting) *LoggerZap {
	var level zapcore.Level
	switch cfg.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()

	hook := &lumberjack.Logger{
		Filename:   getLogFileName(),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	writer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(hook),
	)

	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(core, zap.AddCaller())

	return &LoggerZap{logger}
}

func getLogFileName() string {
	return fmt.Sprintf("logs/app-%s.log", time.Now().Format("2006-01-02"))
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

package core

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

func InitLogger(level string) (*zap.SugaredLogger, error) {

	logLevel, err := zapcore.ParseLevel(level)

	if err != nil {
		panic(err)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConf)
	fileCore := zapcore.NewCore(encoder, writeSyncer, logLevel)
	consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), logLevel)

	core := zapcore.NewTee(
		fileCore,
		consoleCore,
	)

	Log = zap.New(core, zap.AddCaller()).Sugar()

	return Log, nil
}

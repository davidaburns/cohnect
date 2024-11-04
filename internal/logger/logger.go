package logger

import (
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogType int

const (
	ConsoleLog LogType = iota
	JsonLog
	FileLog
)

func CreateNew(logLevel string, logType LogType) *zap.SugaredLogger {
	level := zapcore.InfoLevel
	switch strings.ToLower(logLevel) {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		log.Fatalf("Invalid log level: %s, please use DEBUG, INFO, WARN, or ERROR levels", strings.ToLower(logLevel))
	}

	var encoder zapcore.Encoder
	switch logType {
	case ConsoleLog:
		encoder = initConsoleEncoder()
	case JsonLog:
		encoder = initJsonEncoder()
	default:
		log.Fatalf("Invalid log type: %d", logType)
	}

	core := zapcore.NewCore(encoder, os.Stdout, level)
	options := []zap.Option{zap.AddCaller()}

	return zap.New(core, options...).Sugar()
}


func initConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "lvl",
		NameKey: "log",
		CallerKey: "call",
		MessageKey: "msg",
		StacktraceKey: "stack",
		EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})
}

func initJsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "lvl",
		NameKey: "log",
		CallerKey: "call",
		MessageKey: "msg",
		StacktraceKey: "stack",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})
}
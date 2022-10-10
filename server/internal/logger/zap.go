// Package logger contains logging tools.
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZap provides logging with zap.
func InitZap() {
	logLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// Standard output
	stdCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config()),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	// Output as a log file
	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config()),
		zapcore.AddSync(setFile()),
		logLevel,
	)

	logger := zap.New(zapcore.NewTee(
		logCore,
	))

	dm, ok := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	if ok != nil {
		fmt.Printf("load .env file error: %v\n", ok)
	}

	if dm {
		logger = zap.New(zapcore.NewTee(
			stdCore,
			logCore,
		))
	}

	zap.ReplaceGlobals(logger)
}

// setFile return the location where the log file will be placed.
func setFile() (f *os.File) {
	dirPath := "/var/log"
	fileName := "server.json"
	content := filepath.Join(dirPath, fileName)

	if _, err := os.Stat(content); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(content); err != nil {
				fmt.Println(err)
			}
		}
	}

	f, err := os.OpenFile(content, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

// config returns EncoderConfig for production environments.
func config() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	cfg.MessageKey = "msg"
	cfg.LevelKey = "level"
	cfg.NameKey = "name"
	cfg.TimeKey = "timestamp"
	cfg.CallerKey = "caller"
	cfg.FunctionKey = "func"
	cfg.StacktraceKey = "stacktrace"
	cfg.LineEnding = "\n"
	cfg.EncodeTime = zapcore.EpochTimeEncoder
	cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncodeCaller = zapcore.ShortCallerEncoder

	return cfg
}

// LogDebug is Key-value format debug log.
func LogDebug(msg string, kv ...interface{}) {
	zap.S().Debugw(msg, kv...)
}

// LogErr is Key-value format error log.
func LogErr(msg string, kv ...interface{}) {
	zap.S().Errorw(msg, kv...)
}

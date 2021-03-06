package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	var err error

	cfg := zap.NewProductionConfig()

	log, err = cfg.Build(zap.AddCallerSkip(1))
	cfg.OutputPaths = []string{os.Getenv("LOGGER_DIR")}

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func InfoHttpRequest(r *http.Request) {
	msg := fmt.Sprintf("%s - %s", r.Method, r.URL.EscapedPath())
	var requestBody interface{}
	json.NewDecoder(r.Body).Decode(&requestBody)

	log.Info(
		msg,
		zap.String("method", r.Method),
		zap.String("url", r.URL.EscapedPath()),
		zap.Any("query", r.URL.Query()),
		zap.Any("body", requestBody),
	)
}

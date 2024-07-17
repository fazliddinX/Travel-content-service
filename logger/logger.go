package logger

import (
	"log"
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logFile, err := os.OpenFile("logger/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, nil)
	return slog.New(handler)
}

package logger

import (
	"log/slog"
	"os"
)

// I will visit this back later

func SetTextLogger() {
	slogger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	// slogger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	// 	Level:     slog.LevelDebug,
	// 	AddSource: true,
	// }))
	slog.SetDefault(slogger)
}
func SetJsonLogger() {
	slogger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// slogger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
	// 	Level:     slog.LevelDebug,
	// 	AddSource: true,
	// }))
	slog.SetDefault(slogger)
}

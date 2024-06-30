package main

import (
	"log/slog"
	"os"

	"github.com/haapjari/nta-go/pkg/ui"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	slog.Info("Network Traffic Analyzer")

	userInterface, err := ui.New()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = userInterface.Serve()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	defer func() {
		err = userInterface.Close()
		if err != nil {
			slog.Error(err.Error())
			return
		}
	}()
}

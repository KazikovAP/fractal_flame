package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/KazikovAP/fractal_flame/config"
	"github.com/KazikovAP/fractal_flame/internal/application"
	ts "github.com/KazikovAP/fractal_flame/internal/domain/transformations"
	"github.com/KazikovAP/fractal_flame/internal/infrastructure"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Ошибка инициализации файла конфигурации: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ioAdapter := infrastructure.NewIOAdapter(os.Stdout, logger)

	factory, err := ts.NewTransformationFactory(cfg.TransformFn, ts.RandomColor)
	if err != nil {
		log.Fatalf("Ошибка создания фабрики трансформаций: %v", err)
	}

	app := application.NewApp(cfg, ioAdapter, factory)
	if err := app.Start(); err != nil {
		logger.Error("Application failed to start", "error", err)
	}
}

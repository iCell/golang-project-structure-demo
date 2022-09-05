package app

import "demo/internal/pkg/service"

type App struct {
	Device *service.DeviceUseCase
	// ... etc
	// log, metric
}

func NewApp(deviceUseCase *service.DeviceUseCase) *App {
	return &App{deviceUseCase}
}

package main

import (
	"duepe/config"
	"duepe/config/types"
	"duepe/internal/app/httpsrv"
	"go.uber.org/fx"
	"log/slog"
)

func main() {
	slog.Info("Start Application")

	cfg, err := config.LoadCMDWeb()
	if err != nil {
		slog.Error("error load config: ", slog.String("err", err.Error()))

		return
	}

	httpApp := httpsrv.NewApp(provideParams(cfg))

	httpApp.Run()
	if httpApp.Err() != nil {
		slog.Error("App start error: ", slog.String("err", httpApp.Err().Error()))

		return
	}

	slog.Info("Application finished")
}

func provideParams(cfg *types.CMDWeb) fx.Option {
	return fx.Provide(
		func() *types.Database { return cfg.Database },
		func() *types.HTTPServer { return cfg.HTTPServer },
		func() *types.WebFs { return cfg.WebFs },
	)
}

package httpsrv

import (
	"context"
	"duepe/config/types"
	"duepe/internal/lib/fxutils"
	"duepe/internal/lib/util"
	"errors"
	"log/slog"
	"net/http"
)

type Application struct {
	srv    *http.Server
	tls    *types.HTTPServerTLS
	runner *fxutils.Runner
}

func newApplication(cfg *types.HTTPServer, router *http.ServeMux, runner *fxutils.Runner) *Application {
	srv := &http.Server{
		Addr:         *cfg.Host + ":" + *cfg.Port,
		IdleTimeout:  util.ToTimeDurationSec(*cfg.IdleTimeoutSec),
		WriteTimeout: util.ToTimeDurationSec(*cfg.WriteTimeoutSec),
		ReadTimeout:  util.ToTimeDurationSec(*cfg.ReadTimeoutSec),

		Handler: router,
	}

	return &Application{
		srv:    srv,
		tls:    cfg.TLS,
		runner: runner,
	}
}

func (app *Application) Run(_ context.Context) error {
	go app.serve()
	app.runner.StartTracking(appName)

	slog.Info("HTTP server started", slog.String("addr", app.srv.Addr))

	return nil
}

func (app *Application) Stop(ctx context.Context) error {
	if err := app.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil

}

func (app *Application) serve() {
	if err := app.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("HTTP server failed to start", slog.String("error", err.Error()))
	}

	app.runner.StopTracking(appName)
}

package httpsrv

import (
	"context"
	"duepe/internal/app/httpsrv/static"
	"duepe/internal/lib/fxutils"
	"go.uber.org/fx"
)

const appName = "HTTP-Server"

func NewApp(params fx.Option) *fxutils.App {
	return fxutils.NewApp(
		NewModule(params),
	)
}

func NewModule(params fx.Option) fx.Option {
	return fx.Module(
		appName,
		params,
		
		fx.Options(
			static.NewModule(),
		),

		baseProvide(),
		baseInvoke(),
	)
}

func baseInvoke() fx.Option {
	return fx.Invoke(
		func(lc fx.Lifecycle, app *Application, runner *fxutils.Runner) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return app.Run(ctx)
				},
				OnStop: func(ctx context.Context) error {
					return app.Stop(ctx)
				},
			})
		},
	)
}

func baseProvide() fx.Option {
	return fx.Provide(
		newApplication,
	)
}

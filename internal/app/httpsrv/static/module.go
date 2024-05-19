package static

import (
	"duepe/internal/app/httpsrv/static/servprovider"
	"go.uber.org/fx"
)

const moduleName = "static"

func NewModule() fx.Option {
	return fx.Module(
		moduleName,

		servprovider.ProvideRouter(),
	)
}

package servprovider

import (
	"duepe/internal/app/httpsrv/static/handlers"
	"duepe/internal/app/httpsrv/static/middlewares"
	"duepe/web"
	"go.uber.org/fx"
	"net/http"
)

func ProvideRouter() fx.Option {
	return fx.Options(
		fx.Provide(provideInstanceRouter),
		fx.Invoke(provideIndexRoutes),
		fx.Invoke(provideAssetsRoutes),
	)
}

func provideInstanceRouter() *http.ServeMux {
	return http.NewServeMux()
}

func provideIndexRoutes(router *http.ServeMux) {
	router.HandleFunc("/", handlers.Index)
}

func provideAssetsRoutes(router *http.ServeMux) {
	assetMW := middlewares.NewAsset()
	router.Handle(
		"/assets/",
		assetMW.Apply(
			http.FileServer(http.FS(web.GetPublicFs())),
		),
	)
}

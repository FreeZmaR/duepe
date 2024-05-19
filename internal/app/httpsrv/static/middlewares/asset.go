package middlewares

import (
	"net/http"
	"strings"
)

type Asset struct{}

func NewAsset() *Asset {
	return &Asset{}
}

func (a *Asset) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ReplaceAll(r.URL.Path, "/assets", "/static/assets")

		next.ServeHTTP(w, r)
	})
}

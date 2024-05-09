package web

import "embed"

//go:embed static/assets/*
//go:embed static/index.html
var public embed.FS

// GetPublicFs returns the embedded filesystem.
func GetPublicFs() embed.FS {
	return public
}

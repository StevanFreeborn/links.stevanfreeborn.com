// Package assets contains embedded static assets for the application.
package assets

import (
	"embed"
)

var (
	//go:embed "templates/*"
	Templates embed.FS
	//go:embed "css/*"
	CSS embed.FS
	//go:embed "fonts/*"
	Fonts embed.FS
	//go:embed "json/*"
	JSON embed.FS
	//go:embed "images/*"
	Images embed.FS
)

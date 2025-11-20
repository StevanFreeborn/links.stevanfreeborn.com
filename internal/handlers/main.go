// Package handlers provides HTTP handlers for the web application.
package handlers

import (
	"net/http"
	"text/template"

	"github.com/StevanFreeborn/links.stevanfreeborn.com/internal/assets"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Links coming in external JSON data
	// { link: "link", text: "text", "icon": "icon" }

	t, err := template.ParseFS(assets.Templates, "templates/index.gohtml")

	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func CSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.CSS, r.URL.Path)
}

func Fonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.Fonts, r.URL.Path)
}

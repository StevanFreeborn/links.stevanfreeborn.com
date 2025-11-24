// Package handlers provides HTTP handlers for the web application.
package handlers

import (
	"math"
	"net/http"
	"text/template"
	"time"

	"github.com/StevanFreeborn/links.stevanfreeborn.com/internal/assets"
)

const DAYS_IN_YEAR = 365
const HOURS_IN_DAY = 24

var birthday time.Time = time.Date(1993, time.April, 21, 0, 0, 0, 0, time.UTC)

type IndexViewModel struct {
	Age float64
}

func Index(w http.ResponseWriter, r *http.Request) {
	// TODO: Links coming in external JSON data
	// this could be static file in the repo
	// or it could be fetched
	// { link: "link", text: "text", "icon": "icon" }

	t, err := template.ParseFS(assets.Templates, "templates/index.gohtml")

	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	age := time.Since(birthday).Hours() / HOURS_IN_DAY / DAYS_IN_YEAR

	viewModel := IndexViewModel{
		Age: math.Floor(age),
	}

	t.Execute(w, viewModel)
}

func CSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.CSS, r.URL.Path)
}

func Fonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.Fonts, r.URL.Path)
}

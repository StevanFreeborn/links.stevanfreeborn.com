// Package handlers provides HTTP handlers for the web application.
package handlers

import (
	"encoding/json"
	"io/fs"
	"math"
	"net/http"
	"text/template"
	"time"

	"github.com/StevanFreeborn/links.stevanfreeborn.com/internal/assets"
)

const LINKS_JSON_PATH = "json/links.json"
const DAYS_IN_YEAR = 365
const HOURS_IN_DAY = 24

var birthday time.Time = time.Date(1993, time.April, 21, 0, 0, 0, 0, time.UTC)

type Link struct {
	Href string `json:"href"`
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type IndexViewModel struct {
	Age   float64
	Links []Link
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(assets.Templates, "templates/index.gohtml")

	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	linksFile, err := assets.JSON.Open(LINKS_JSON_PATH)

	if err != nil {
		http.Error(w, "Unable to load links", http.StatusInternalServerError)
		return
	}

	links, err := readLinksFromJSON(linksFile)

	if err != nil {
		http.Error(w, "Unable to parse links", http.StatusInternalServerError)
		return
	}

	age := time.Since(birthday).Hours() / HOURS_IN_DAY / DAYS_IN_YEAR

	viewModel := IndexViewModel{
		Age:   math.Floor(age),
		Links: links,
	}

	t.Execute(w, viewModel)
}

func CSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.CSS, r.URL.Path)
}

func Fonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.Fonts, r.URL.Path)
}

func Images(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, assets.Images, r.URL.Path)
}

func readLinksFromJSON(file fs.File) ([]Link, error) {
	defer file.Close()
	var links []Link

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&links)

	if err != nil {
		return nil, err
	}

	return links, nil
}

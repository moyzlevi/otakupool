package main

import (
	"encoding/base64"
	"html/template"
	"path/filepath"

	"github.com/moyzlevi/otakupool/internal/models"
)

type templateData struct {
	Images []*models.Image
}

func byteArrayToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

var functions = template.FuncMap {
	"byteArrayToBase64": byteArrayToBase64,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
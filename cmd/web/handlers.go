package main

import (
	"io"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) poolView(w http.ResponseWriter, r *http.Request) {
	images, err := app.images.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Images = images
	app.render(w, http.StatusOK, "pool-view.tmpl.html", data)
}

func (app *application) createGet(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("uploadedFile")
	if err != nil {
		app.serverError(w, err)
        return
    }
    defer file.Close()

	fileBytes, err := io.ReadAll(file)
    if err != nil {
        app.serverError(w, err)
        return
    }

	app.images.Insert(fileBytes, handler.Header.Get("Content-Type"), handler.Filename)
	http.Redirect(w, r, "/pool-view" , http.StatusSeeOther)
}

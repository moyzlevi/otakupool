package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/moyzlevi/otakupool/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	templateCache map[string]*template.Template
	images *models.ImageModel
}

func main() {
	cfg := initConf()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDb(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()

	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		templateCache: templateCache,
		images: &models.ImageModel{DB: db},
	}

	srv := &http.Server{
		Addr:         cfg.addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting Server on %s...", cfg.addr)
	errorLog.Fatal(srv.ListenAndServe())
}

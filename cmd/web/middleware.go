package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err:= recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w,r)
	})
}

func secureHeaders(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", 
		"default-src 'self'; style-src 'self' https://fonts.googleapis.com/css2?family=Oxanium:wght@700&display=swap font.googleapis.com;img-src 'self' data:; font-src fonts.gstatic.com")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		next.ServeHTTP(w,r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s",
		r.RemoteAddr,
		r.Proto,
		r.Method,
		r.URL.RequestURI())
	
		next.ServeHTTP(w,r)
	})
}
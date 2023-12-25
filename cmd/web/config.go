package main

import "flag"

type config struct {
	addr string
	staticDir string
	dsn string
}

func initConf() (config) {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir,"staticDir", "./ui/static/", "Path to static folder")
	flag.StringVar(&cfg.dsn,"dsn", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", "The connection string of the database")
	flag.Parse()
	return cfg
}
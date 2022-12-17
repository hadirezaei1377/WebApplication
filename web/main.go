package main

import (
	"log"
	"text/template"
)

const version = "1.0.0"
const cssVersion = "1"

// configuration info for app
type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

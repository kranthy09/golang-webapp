package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	Inproduction  bool
	IngoLog       *log.Logger
	Session       *scs.SessionManager
}

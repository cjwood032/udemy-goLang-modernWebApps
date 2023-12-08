package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type ApplConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}

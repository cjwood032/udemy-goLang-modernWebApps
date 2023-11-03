package config

import "html/template"

type ApplConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}

package config

import "html/template"

type ApplConfig struct {
	TemplateCache map[string]*template.Template
}

package config

import "html/template"

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}

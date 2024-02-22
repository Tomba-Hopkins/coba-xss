package views

import (
	"embed"
	"text/template"
)

//go:embed html
var templates embed.FS
var MyTemplates = template.Must(template.ParseFS(templates, "html/*.html"))
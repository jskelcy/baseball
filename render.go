package main

import (
	"html/template"
	"io"
)

type renderWrapper struct {
	League *league
}

func render(l *league, w io.Writer) {
	wrapper := renderWrapper{League: l}
	tmpl := template.Must(template.ParseFiles("template.html"))
	tmpl.Execute(w, wrapper)
}

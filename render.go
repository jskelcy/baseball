package main

import (
	"html/template"
	"io"
)

type renderWrapper struct {
	League []*fantasyTeam
}

func render(l fantasypctLeague, w io.Writer) {
	wrapper := renderWrapper{League: l.GetTeams()}
	tmpl := template.Must(template.ParseFiles("template.html"))
	tmpl.Execute(w, wrapper)
}

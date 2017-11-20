/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

// Package {{.Package}} provides commands and events for {{.Domain}} domain.
package {{.Package}}

import (
    "github.com/SKatiyar/cri"
{{- range .Imports}}
    {{.Name}} {{printf "%q" .Path -}}
{{end}}
)

// List of commands in {{.Domain}} domain
const (
{{- range .Commands}}
    {{.Name}} = {{printf "%q" .Command -}}
{{end}}
)

{{- if ne (len .Events) 0}}
// List of events in {{.Domain}} domain
const (
{{- range .Events}}
    {{.Name}} = {{printf "%q" .Command -}}
{{end}}
)
{{end}}

{{.Doc}}
type {{.Domain}} struct {
	conn cri.Connector
}

// New creates a {{.Domain}} instance
func New(conn cri.Connector) *{{.Domain}} {
	return &{{.Domain}}{conn}
}

{{- range .Commands}}
{{- range .Types}}
{{.Doc}}
type {{.ID}} {{.Type}}
{{end}}
{{.Doc}}
func (obj *{{.Domain}}) {{.Name}}({{.RequestName}}) ({{.ResponseName}}) {
	err = obj.conn.Send({{.Name}}, {{.RequestValue}}, {{.ResponseValue}})
	return
}
{{end -}}

{{range .Events}}
{{- range .Types}}
{{.Doc}}
type {{.ID}} {{.Type}}
{{end}}
{{.Doc}}
func (obj *{{.Domain}}) {{.Name}}(fn func({{.EventParams}}) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On({{.Name}}, closeChn)
	go func() {
		for {
            {{.ParamsDecl}}
			readErr := decoder({{.ParamsValue}})
            if !fn({{.CallParams}}) {
				close(closeChn)
				break
			}
		}
	}()
}
{{end -}}

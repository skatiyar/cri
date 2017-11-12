/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

{{.Doc}}
package {{.Package}}

import (
    "github.com/SKatiyar/cri"
{{- range .Imports}}
    {{.Name}} {{printf "%q" .Path -}}
{{end}}
)

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
	err = obj.conn.Send({{printf "%q" .Command}}, {{.RequestValue}}, {{.ResponseValue}})
	return
}
{{end -}}

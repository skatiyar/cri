/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package {{.Package}}

{{range .Types}}
{{.Doc}}
type {{.ID}} {{.Type}}
{{end}}

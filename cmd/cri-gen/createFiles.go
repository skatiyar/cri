package main

import (
	"path"
	"strings"
	"text/template"

	"go/ast"
	"go/token"
)

const ImportPath = "github.com/SKatiyar/cri"

func parseBasicType(typ string) ast.Expr {
	if typ == "integer" {
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: "int",
		}
	} else if typ == "number" {
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: "float32",
		}
	} else if typ == "string" {
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: "string",
		}
	} else if typ == "boolean" {
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: "bool",
		}
	} else if typ == "object" {
		return &ast.MapType{
			Key: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "string",
			},
			Value: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "interface{}",
			},
		}
	} else if typ == "any" {
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: "interface{}",
		}
	} else {
		panic("Unhandled type " + typ)
	}
}

func parseRef(ref, domain, structName string, notTypes bool) ast.Expr {
	if idx := strings.Index(ref, "."); idx != -1 {
		if notTypes {
			return &ast.BasicLit{
				Kind:  token.STRING,
				Value: "types." + strings.Replace(ref, ".", "_", -1),
			}
		} else {
			return &ast.BasicLit{
				Kind:  token.STRING,
				Value: strings.Replace(ref, ".", "_", -1),
			}
		}
	} else {
		if ref == structName {
			return &ast.StarExpr{
				X: &ast.BasicLit{
					Kind:  token.STRING,
					Value: domain + "_" + ref,
				},
			}
		} else if notTypes {
			return &ast.BasicLit{
				Kind:  token.STRING,
				Value: "types." + domain + "_" + ref,
			}
		} else {
			return &ast.BasicLit{
				Kind:  token.STRING,
				Value: domain + "_" + ref,
			}
		}
	}
}

func parseParameter(p Parameter, domain, structName string, notTypes bool) (*ast.Field, []string) {
	deps := make([]string, 0)
	nField := &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(strings.Title(p.Name)),
		},
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{},
		},
	}

	if len(p.Description) > 0 {
		nField.Doc.List = append(nField.Doc.List, &ast.Comment{
			Text: "\n// " + p.Description,
		})
	}
	if p.Experimental {
		nField.Doc.List = append(nField.Doc.List, &ast.Comment{
			Text: "// NOTE Experimental",
		})
	}

	if p.Optional {
		nField.Tag = &ast.BasicLit{
			Kind:  token.STRING,
			Value: "`json:\"" + p.Name + ",omitempty\"`",
		}
	} else {
		nField.Tag = &ast.BasicLit{
			Kind:  token.STRING,
			Value: "`json:\"" + p.Name + "\"`",
		}
	}

	if len(p.Type) != 0 {
		if p.Type == "array" {
			if len(p.Items.Type) > 0 {
				nField.Type = &ast.ArrayType{
					Elt: parseBasicType(p.Items.Type),
				}
			} else if len(p.Items.Ref) > 0 {
				nField.Type = &ast.ArrayType{
					Elt: parseRef(p.Items.Ref, domain, structName, notTypes),
				}
				if idx := strings.Index(p.Items.Ref, "."); idx != -1 {
					deps = append(deps, string(p.Items.Ref[:idx]))
				} else if notTypes {
					deps = append(deps, "types")
				}
			} else {
				nField.Type = &ast.ArrayType{
					Elt: &ast.BasicLit{
						Kind:  token.STRING,
						Value: "interface{}",
					},
				}
			}
		} else if p.Type == "integer" ||
			p.Type == "number" ||
			p.Type == "any" ||
			p.Type == "object" ||
			p.Type == "boolean" ||
			p.Type == "string" {
			if p.Optional && p.Type != "any" {
				nField.Type = &ast.StarExpr{
					X: parseBasicType(p.Type),
				}
			} else {
				nField.Type = parseBasicType(p.Type)
			}
		} else {
			panic("Unhandled type " + p.Type)
		}
	} else if len(p.Ref) != 0 {
		if p.Optional && p.Ref != structName {
			nField.Type = &ast.StarExpr{
				X: parseRef(p.Ref, domain, structName, notTypes),
			}
		} else {
			nField.Type = parseRef(p.Ref, domain, structName, notTypes)
		}
		if idx := strings.Index(p.Ref, "."); idx != -1 {
			deps = append(deps, string(p.Ref[:idx]))
		} else if notTypes {
			deps = append(deps, "type")
		}
	}

	return nField, deps
}

type TypeData struct {
	ID   string
	Doc  string
	Type string
}

type TypesData struct {
	Package string
	Types   []TypeData
}

func transformBasicTypes(typ string) string {
	if typ == "integer" {
		return "int"
	} else if typ == "number" {
		return "float32"
	} else if typ == "string" {
		return "string"
	} else if typ == "boolean" {
		return "bool"
	} else if typ == "object" {
		return "map[string]interface{}"
	} else if typ == "any" {
		return "interface{}"
	} else {
		panic("Unhandled type " + typ)
	}
}

func transformParameter(p Parameter, domain, structName string, notTypes bool) (string, []string) {
	nField := make([]string, 0)
	deps := make([]string, 0)

	if len(p.Description) > 0 {
		nField = append(nField, "\t// "+strings.Trim(p.Description, "\n"))
	}
	if p.Experimental {
		nField = append(nField, "\t// NOTE Experimental")
	}

	fieldSpecs := []string{"\t" + strings.Title(p.Name)}
	if len(p.Type) != 0 {
		if p.Type == "array" {
			if len(p.Items.Type) > 0 {
				fieldSpecs = append(fieldSpecs, "[]"+transformBasicTypes(p.Items.Type))
			} else if len(p.Items.Ref) > 0 {
				stm, stmDeps := transformRef(p.Items.Ref, domain, structName, notTypes)
				fieldSpecs = append(fieldSpecs, "[]"+stm)
				deps = append(deps, stmDeps...)
			} else {
				fieldSpecs = append(fieldSpecs, "[]interface{}")
			}
		} else if p.Type == "integer" ||
			p.Type == "number" ||
			p.Type == "boolean" ||
			p.Type == "any" ||
			p.Type == "object" ||
			p.Type == "string" {
			if p.Optional && p.Type != "any" && p.Type != "object" {
				fieldSpecs = append(fieldSpecs, "*"+transformBasicTypes(p.Type))
			} else {
				fieldSpecs = append(fieldSpecs, transformBasicTypes(p.Type))
			}
		} else {
			panic("Unhandled type " + p.Type)
		}
	} else if len(p.Ref) != 0 {
		if p.Optional && p.Ref != structName {
			stm, stmDeps := transformRef(p.Ref, domain, structName, notTypes)
			fieldSpecs = append(fieldSpecs, "*"+stm)
			deps = append(deps, stmDeps...)
		} else {
			stm, stmDeps := transformRef(p.Ref, domain, structName, notTypes)
			fieldSpecs = append(fieldSpecs, stm)
			deps = append(deps, stmDeps...)
		}
	}

	if p.Optional {
		fieldSpecs = append(fieldSpecs, "`json:\""+p.Name+",omitempty\"`")
	} else {
		fieldSpecs = append(fieldSpecs, "`json:\""+p.Name+"\"`")
	}

	nField = append(nField, strings.Join(fieldSpecs, " "))

	return strings.Join(nField, "\n"), deps
}

func transformRef(ref, domain, structName string, notTypes bool) (string, []string) {
	if idx := strings.Index(ref, "."); idx != -1 {
		if notTypes {
			return "types." + strings.Replace(ref, ".", "_", -1), []string{domain}
		} else {
			return strings.Replace(ref, ".", "_", -1), []string{string(ref[:idx])}
		}
	} else {
		if ref == structName {
			return "*" + domain + "_" + ref, nil
		} else if notTypes {
			return "types." + domain + "_" + ref, []string{domain}
		} else {
			return domain + "_" + ref, []string{domain}
		}
	}
}

func transformTypes(d Domain) TypesData {
	data := TypesData{
		Package: "types",
	}

	imports := make(map[string]bool)
	typs := d.Types
	for i := 0; i < len(typs); i++ {
		typeSpec := TypeData{
			ID: d.Domain + "_" + typs[i].ID,
		}
		if len(typs[i].Description) > 0 {
			typeSpec.Doc = "//" + typs[i].Description
		}

		if len(typs[i].Enum) != 0 {
			// TODO Add enums to documentation
			// NOTE Consider using enums pattern
			// for defining as in
			// http://golang-basic.blogspot.in/2014/07/step-by-step-guide-to-declaring-enums.html
		}

		if typs[i].Type == "array" {
			if len(typs[i].Items.Type) > 0 {
				typeSpec.Type = "[]" + transformBasicTypes(typs[i].Items.Type)
			} else if len(typs[i].Items.Ref) > 0 {
				stm, stmDeps := transformRef(typs[i].Items.Ref, d.Domain, typs[i].ID, false)
				typeSpec.Type = "[]" + stm
				for _, dep := range stmDeps {
					imports[dep] = true
				}
			} else {
				typeSpec.Type = "[]interface{}"
			}
		} else if typs[i].Type == "integer" ||
			typs[i].Type == "number" ||
			typs[i].Type == "boolean" ||
			typs[i].Type == "any" ||
			typs[i].Type == "string" {
			typeSpec.Type = transformBasicTypes(typs[i].Type)
		} else if typs[i].Type == "object" {
			properties := make([]string, 0)
			for j := 0; j < len(typs[i].Properties); j++ {
				param, paramDeps := transformParameter(typs[i].Properties[j], d.Domain, typs[i].ID, false)
				properties = append(properties, param)
				for _, dep := range paramDeps {
					imports[dep] = true
				}
			}
			typeSpec.Type = "struct {\n" + strings.Join(properties, "\n") + "\n}"
		} else {
			panic("Unhandled type " + typs[i].Type)
		}

		data.Types = append(data.Types, typeSpec)
	}

	return data
}

type CommandImport struct {
	Name string
	Path string
}

type CommandData struct {
	Doc           string
	Domain        string
	Command       string
	Name          string
	RequestName   string
	RequestValue  string
	ResponseName  string
	ResponseValue string
	Types         []TypeData
}

type CommandsData struct {
	Doc      string
	Domain   string
	Package  string
	Imports  []CommandImport
	Commands []CommandData
}

func transformCommands(d Domain) CommandsData {
	data := CommandsData{
		Domain:   d.Domain,
		Package:  strings.ToLower(d.Domain),
		Imports:  make([]CommandImport, 0),
		Commands: make([]CommandData, 0),
	}

	if len(d.Description) > 0 {
		data.Doc = "// " + d.Description
	}

	imports := make(map[string]bool)
	for i := 0; i < len(d.Commands); i++ {
		cmd := CommandData{
			Domain:  d.Domain,
			Name:    strings.Title(d.Commands[i].Name),
			Command: d.Domain + "." + d.Commands[i].Name,
			Types:   make([]TypeData, 0),
		}
		if len(d.Commands[i].Description) > 0 {
			cmd.Doc = "// " + d.Commands[i].Description
		}
		if len(d.Commands[i].Parameters) > 0 {
			reqName := strings.Title(d.Commands[i].Name) + "Request"
			cmd.RequestName = "request *" + reqName
			cmd.RequestValue = "request"
			properties := make([]string, 0)
			for j := 0; j < len(d.Commands[i].Parameters); j++ {
				param, paramDeps := transformParameter(d.Commands[i].Parameters[j], d.Domain, reqName, true)
				properties = append(properties, param)
				for _, dep := range paramDeps {
					imports[dep] = true
				}
			}
			cmd.Types = append(cmd.Types, TypeData{
				ID:   reqName,
				Type: "struct {\n" + strings.Join(properties, "\n") + "\n}",
			})
		} else {
			cmd.RequestName = ""
			cmd.RequestValue = "nil"
		}
		if len(d.Commands[i].Returns) > 0 {
			resName := strings.Title(d.Commands[i].Name) + "Response"
			cmd.ResponseName = "response " + resName + ", err error"
			cmd.ResponseValue = "&response"
			properties := make([]string, 0)
			for j := 0; j < len(d.Commands[i].Returns); j++ {
				param, paramDeps := transformParameter(d.Commands[i].Returns[j], d.Domain, resName, true)
				properties = append(properties, param)
				for _, dep := range paramDeps {
					imports[dep] = true
				}
			}
			cmd.Types = append(cmd.Types, TypeData{
				ID:   resName,
				Type: "struct {\n" + strings.Join(properties, "\n") + "\n}",
			})
		} else {
			cmd.ResponseName = "err error"
			cmd.ResponseValue = "nil"
		}

		data.Commands = append(data.Commands, cmd)
	}
	if len(imports) > 0 {
		data.Imports = append(data.Imports, CommandImport{
			Name: "types",
			Path: path.Join(ImportPath, "types"),
		})
	}

	return data
}

func createTypesFile() (*template.Template, error) {
	return template.ParseFiles(TypesTplFile)
}

func createCommandsFile() (*template.Template, error) {
	return template.ParseFiles(CommandsTplFile)
}

func createVersionFile() (*template.Template, error) {
	return template.ParseFiles(VersionTplFile)
}

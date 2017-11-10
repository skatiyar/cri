package main

import (
	"path"
	"strings"

	"go/ast"
	"go/token"
)

const ImportPath = "github.com/SKatiyar/cri"

func toPackageName(s string) string {
	strs := strings.Split(s, ".")
	strs[0] = strings.ToLower(strs[0])

	return strings.Join(strs, ".")
}

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
		/*
			if idx := strings.Index(p.Ref, "."); idx != -1 {
				nField.Type = &ast.BasicLit{
					Kind:  token.STRING,
					Value: toPackageName(p.Ref),
				}
				if domain != string(p.Ref[:idx]) {
					deps = append(deps, string(p.Ref[:idx]))
				}
			} else {
				if p.Ref == structName {
					nField.Type = &ast.StarExpr{
						X: &ast.BasicLit{
							Kind:  token.STRING,
							Value: p.Ref,
						},
					}
				} else if notTypes {
					nField.Type = &ast.BasicLit{
						Kind:  token.STRING,
						Value: "types." + p.Ref,
					}
					deps = append(deps, domain)
				} else {
					nField.Type = &ast.BasicLit{
						Kind:  token.STRING,
						Value: p.Ref,
					}
				}
			}
		*/
	}

	return nField, deps
}

func parseCommand(c Command, d Domain) (*ast.FuncDecl, []string) {
	resultsFieldList := make([]*ast.Field, 0)
	callParams := make([]ast.Expr, 0)
	deps := make([]string, 0)

	for i := 0; i < len(c.Returns); i++ {
		param, ldeps := parseParameter(c.Returns[i], d.Domain, "", true)
		resultsFieldList = append(resultsFieldList, param)
		deps = append(deps, ldeps...)
	}

	callParams = append(callParams, &ast.BasicLit{
		Kind:  token.STRING,
		Value: "\"" + d.Domain + "." + c.Name + "\"",
	})
	if len(c.Parameters) > 0 {
		callParams = append(callParams, &ast.BasicLit{
			Kind:  token.STRING,
			Value: "request",
		})
	} else {
		callParams = append(callParams, &ast.BasicLit{
			Kind:  token.STRING,
			Value: "nil",
		})
	}
	if len(c.Returns) > 0 {
		callParams = append(callParams, &ast.BasicLit{
			Kind:  token.STRING,
			Value: "&response",
		})
	} else {
		callParams = append(callParams, &ast.BasicLit{
			Kind:  token.STRING,
			Value: "nil",
		})
	}

	nFunc := &ast.FuncDecl{
		Name: ast.NewIdent(strings.Title(c.Name)),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent("obj"),
					},
					Type: &ast.StarExpr{
						X: &ast.BasicLit{
							Kind:  token.STRING,
							Value: strings.Title(d.Domain),
						},
					},
				},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{
							ast.NewIdent("err"),
						},
						Type: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "error",
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.BinaryExpr{
						X: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "err",
						},
						Op: token.ASSIGN,
						Y: &ast.CallExpr{
							Fun: &ast.BasicLit{
								Kind:  token.STRING,
								Value: "obj.conn.Send",
							},
							Args: callParams,
						},
					},
				},
				&ast.ReturnStmt{},
			},
		},
	}

	if len(c.Parameters) > 0 {
		nFunc.Type.Params.List = append(nFunc.Type.Params.List, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent("request"),
			},
			Type: &ast.StarExpr{
				X: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strings.Title(c.Name) + "Request",
				},
			},
		})
	}
	if len(c.Returns) > 0 {
		nFunc.Type.Results.List = append([]*ast.Field{
			&ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("response"),
				},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: resultsFieldList,
					},
				},
			},
		}, nFunc.Type.Results.List...)
	}

	return nFunc, deps
}

func createVersionFile(v Version, pkgPath string) *ast.File {
	return &ast.File{
		Name: ast.NewIdent(path.Base(pkgPath)),
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("Version"),
				Type: &ast.FuncType{
					Results: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent("major"),
									ast.NewIdent("minor"),
								},
								Type: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "string",
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: "\"" + v.Major + "\"",
								},
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: "\"" + v.Minor + "\"",
								},
							},
						},
					},
				},
			},
		},
	}
}

func createTypesFile(d Domain) *ast.File {
	nFile := &ast.File{
		Name: ast.NewIdent("types"),
	}
	typs := d.Types
	decls := make([]ast.Decl, 0)

	for i := 0; i < len(typs); i++ {
		typeSpec := &ast.TypeSpec{
			Name: ast.NewIdent(d.Domain + "_" + typs[i].ID),
		}

		if len(typs[i].Enum) != 0 {
			// TODO Add enums to documentation
			// NOTE Consider using enums pattern
			// for defining as in
			// http://golang-basic.blogspot.in/2014/07/step-by-step-guide-to-declaring-enums.html
		}

		if typs[i].Type == "array" {
			if len(typs[i].Items.Type) > 0 {
				typeSpec.Type = &ast.ArrayType{
					Elt: parseBasicType(typs[i].Items.Type),
				}
			} else if len(typs[i].Items.Ref) > 0 {
				typeSpec.Type = &ast.ArrayType{
					Elt: parseRef(typs[i].Items.Ref, d.Domain, typs[i].ID, false),
				}
			} else {
				typeSpec.Type = &ast.ArrayType{
					Elt: &ast.BasicLit{
						Kind:  token.STRING,
						Value: "interface{}",
					},
				}
			}
		} else if typs[i].Type == "integer" ||
			typs[i].Type == "number" ||
			typs[i].Type == "boolean" ||
			typs[i].Type == "any" ||
			typs[i].Type == "string" {
			typeSpec.Type = parseBasicType(typs[i].Type)
		} else if typs[i].Type == "object" {
			properties := make([]*ast.Field, 0)
			for j := 0; j < len(typs[i].Properties); j++ {
				param, _ := parseParameter(typs[i].Properties[j], d.Domain, typs[i].ID, false)
				properties = append(properties, param)
			}
			typeSpec.Type = &ast.StructType{
				Fields: &ast.FieldList{
					List: properties,
				},
			}
		} else {
			panic("Unhandled type " + typs[i].Type)
		}

		decls = append(decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				typeSpec,
			},
		})
	}

	nFile.Decls = decls

	return nFile
}

func createCommandsFile(d Domain) *ast.File {
	nFile := &ast.File{
		Name: ast.NewIdent(strings.ToLower(d.Domain)),
	}

	decls := make([]ast.Decl, 0)
	imports := make([]string, 0)
	structName := d.Domain

	structDecl := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(structName),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent("conn"),
								},
								Type: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "cri.Connector",
								},
							},
						},
					},
				},
			},
		},
	}
	decls = append(decls, structDecl)

	constructorDecl := &ast.FuncDecl{
		Name: ast.NewIdent("New"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{
							ast.NewIdent("conn"),
						},
						Type: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "cri.Connector",
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Type: &ast.StarExpr{
							X: &ast.BasicLit{
								Kind:  token.STRING,
								Value: structName,
							},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.UnaryExpr{
							Op: token.AND,
							X: &ast.CompositeLit{
								Type: &ast.BasicLit{
									Kind:  token.STRING,
									Value: structName,
								},
								Elts: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: "conn",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	decls = append(decls, constructorDecl)

	for i := 0; i < len(d.Commands); i++ {
		if len(d.Commands[i].Parameters) > 0 {
			paramsDecls := make([]*ast.Field, 0)
			for j := 0; j < len(d.Commands[i].Parameters); j++ {
				param, ldeps := parseParameter(d.Commands[i].Parameters[j], d.Domain, strings.Title(d.Commands[i].Name)+"Request", true)
				paramsDecls = append(paramsDecls, param)
				imports = append(imports, ldeps...)
			}
			responseDecl := &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent(strings.Title(d.Commands[i].Name) + "Request"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: paramsDecls,
							},
						},
					},
				},
			}
			decls = append(decls, responseDecl)
		}
		cmdDecls, cmdDeps := parseCommand(d.Commands[i], d)
		strings.ToLower(d.Domain)
		decls = append(decls, cmdDecls)
		imports = append(imports, cmdDeps...)
	}

	decls = append([]ast.Decl{
		&ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.IMPORT,
						Value: "\"" + path.Join(ImportPath) + "\"",
					},
				},
			},
		},
	}, decls...)
	if len(imports) > 0 {
		decls = append([]ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Name: ast.NewIdent("types"),
						Path: &ast.BasicLit{
							Kind:  token.IMPORT,
							Value: "\"" + path.Join(ImportPath, "types") + "\"",
						},
					},
				},
			},
		}, decls...)
	}

	nFile.Decls = decls

	return nFile
}

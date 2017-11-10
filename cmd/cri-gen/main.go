package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"go/ast"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
)

const (
	ProtocolOutputPath = "../../"
	FilePerm           = 0700

	JSProtocolFile      = "../../json-proto/js_protocol.json"
	BrowserProtocolFile = "../../json-proto/browser_protocol.json"
)

var (
	ErrMajorVersionMismatch = errors.New("major version of protocols don't match")
	ErrMinorVersionMismatch = errors.New("minor version of protocols don't match")
)

func main() {
	jsProtoData, jsFileErr := ioutil.ReadFile(JSProtocolFile)
	if jsFileErr != nil {
		panic(jsFileErr)
	}

	browserProtoData, browserFileErr := ioutil.ReadFile(BrowserProtocolFile)
	if browserFileErr != nil {
		panic(browserFileErr)
	}

	var jsProto JSProtocol
	if decodeErr := json.Unmarshal(jsProtoData, &jsProto); decodeErr != nil {
		panic(decodeErr)
	}

	var browserProto BrowserProtocol
	if decodeErr := json.Unmarshal(browserProtoData, &browserProto); decodeErr != nil {
		panic(decodeErr)
	}

	createCode(jsProto, browserProto)
}

func createCode(jsProto JSProtocol, browserProto BrowserProtocol) {
	if jsProto.Version.Major != browserProto.Version.Major {
		panic(ErrMajorVersionMismatch)
	}
	if jsProto.Version.Minor != browserProto.Version.Minor {
		panic(ErrMinorVersionMismatch)
	}

	config := types.Config{Importer: importer.For("source", nil)}
	domains := append(jsProto.Domains, browserProto.Domains...)
	commandsFileSet := token.NewFileSet()
	commandsFiles := make([]string, 0)
	typesFileSet := token.NewFileSet()
	typesFiles := make([]string, 0)
	for i := 0; i < len(domains); i++ {
		typesPath := path.Join(ProtocolOutputPath, "types")
		if dirErr := os.Mkdir(typesPath, FilePerm); dirErr != nil {
			if !os.IsExist(dirErr) {
				panic(dirErr)
			}
		}

		commandsPath := path.Join(ProtocolOutputPath, strings.ToLower(domains[i].Domain))
		if dirErr := os.Mkdir(commandsPath, FilePerm); dirErr != nil {
			if !os.IsExist(dirErr) {
				panic(dirErr)
			}
		}

		log.Println("Parsing", domains[i].Domain)

		tFilePath := path.Join(typesPath, strings.ToLower(domains[i].Domain)+".go")
		tFile, tFileErr := os.OpenFile(tFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, FilePerm)
		if tFileErr != nil {
			panic(tFileErr)
		}

		typesFiles = append(typesFiles, tFilePath)
		newTypesFile := createTypesFile(domains[i])
		if writeErr := printer.Fprint(tFile, typesFileSet, newTypesFile); writeErr != nil {
			panic(writeErr)
		}
		if fileCloseErr := tFile.Close(); fileCloseErr != nil {
			panic(fileCloseErr)
		}

		cFilePath := path.Join(ProtocolOutputPath, strings.ToLower(domains[i].Domain), "commands.go")
		cFile, cFileErr := os.OpenFile(cFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, FilePerm)
		if cFileErr != nil {
			panic(cFileErr)
		}

		commandsFiles = append(commandsFiles, cFilePath)
		newCommandsFile := createCommandsFile(domains[i])
		if writeErr := printer.Fprint(cFile, commandsFileSet, newCommandsFile); writeErr != nil {
			panic(writeErr)
		}
		if fileCloseErr := cFile.Close(); fileCloseErr != nil {
			panic(fileCloseErr)
		}
	}

	tFiles := make([]*ast.File, 0)
	for i := 0; i < len(typesFiles); i++ {
		tFile, tFileErr := parser.ParseFile(typesFileSet, typesFiles[i], nil, 0)
		if tFileErr != nil {
			panic(tFileErr)
		}

		tFiles = append(tFiles, tFile)
	}

	log.Println("Checking types")
	tImportPath := path.Join(ImportPath, "cri", "types")
	if _, pkgErr := config.Check(tImportPath, typesFileSet, tFiles, nil); pkgErr != nil {
		panic(pkgErr)
	}

	for i := 0; i < len(commandsFiles); i++ {
		cFile, cFileErr := parser.ParseFile(commandsFileSet, commandsFiles[i], nil, 0)
		if cFileErr != nil {
			panic(cFileErr)
		}

		commandName := path.Base(path.Dir(commandsFiles[i]))
		log.Println("Checking command", commandName)
		cImportPath := path.Join(ImportPath, "cri", commandName)
		if _, pkgErr := config.Check(cImportPath, commandsFileSet, []*ast.File{cFile}, nil); pkgErr != nil {
			panic(pkgErr)
		}
	}

	vFilePath := path.Join(ProtocolOutputPath, "version.go")
	vFile, vFileErr := os.OpenFile(vFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, FilePerm)
	if vFileErr != nil {
		panic(vFileErr)
	}

	vFileSet := token.NewFileSet()
	newVersionFile := createVersionFile(jsProto.Version, "cri")
	if writeErr := printer.Fprint(vFile, vFileSet, newVersionFile); writeErr != nil {
		panic(writeErr)
	}
	if fileCloseErr := vFile.Close(); fileCloseErr != nil {
		panic(fileCloseErr)
	}

	pFile, pFileErr := parser.ParseFile(vFileSet, vFilePath, nil, 0)
	if pFileErr != nil {
		panic(pFileErr)
	}

	pkgImportPath := path.Join(ImportPath, "cri")
	if _, pkgErr := config.Check(pkgImportPath, vFileSet, []*ast.File{pFile}, nil); pkgErr != nil {
		panic(pkgErr)
	}
}

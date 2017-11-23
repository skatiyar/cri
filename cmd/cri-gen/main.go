package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	ProtocolOutputPath = "../../"
	FilePerm           = 0700

	JSProtocolFile      = "json-proto/js_protocol.json"
	BrowserProtocolFile = "json-proto/browser_protocol.json"

	VersionTplFile  = "templates/version.go.tpl"
	CommandsTplFile = "templates/commands.go.tpl"
	TypesTplFile    = "templates/types.go.tpl"
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
	typesPath := path.Join(ProtocolOutputPath, "types")
	if dirErr := os.Mkdir(typesPath, FilePerm); dirErr != nil {
		if !os.IsExist(dirErr) {
			panic(dirErr)
		}
	}

	domains := append(jsProto.Domains, browserProto.Domains...)
	commandsFiles := make([]string, 0)
	typesFiles := make([]string, 0)
	for i := 0; i < len(domains); i++ {
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
		newTypesFile, newTypesFileErr := createTypesFile()
		if newTypesFileErr != nil {
			panic(newTypesFileErr)
		}
		if writeErr := newTypesFile.Execute(tFile, transformTypes(domains[i])); writeErr != nil {
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
		newCommandsFile, newCommandsFileErr := createCommandsFile()
		if newCommandsFileErr != nil {
			panic(newCommandsFileErr)
		}
		if writeErr := newCommandsFile.Execute(cFile, transformCommands(domains[i])); writeErr != nil {
			panic(writeErr)
		}
		if fileCloseErr := cFile.Close(); fileCloseErr != nil {
			panic(fileCloseErr)
		}
	}

	log.Println("Parsing", "Version")

	vFilePath := path.Join(ProtocolOutputPath, "version.go")
	vFile, vFileErr := os.OpenFile(vFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, FilePerm)
	if vFileErr != nil {
		panic(vFileErr)
	}

	newVersionFile, newVersionFileErr := createVersionFile()
	if newVersionFileErr != nil {
		panic(newVersionFileErr)
	}

	if writeErr := newVersionFile.Execute(vFile, struct{ BMajor, BMinor, JMajor, JMinor string }{
		BMajor: browserProto.Version.Major,
		BMinor: browserProto.Version.Minor,
		JMajor: jsProto.Version.Major,
		JMinor: jsProto.Version.Minor,
	}); writeErr != nil {
		panic(writeErr)
	}
	if fileCloseErr := vFile.Close(); fileCloseErr != nil {
		panic(fileCloseErr)
	}
}

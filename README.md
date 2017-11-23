# cri - chrome remote interface

[![GoDoc](https://godoc.org/github.com/SKatiyar/cri?status.svg)](https://godoc.org/github.com/SKatiyar/cri)
[![Go Report Card](https://goreportcard.com/badge/github.com/SKatiyar/cri)](https://goreportcard.com/report/github.com/SKatiyar/cri)
[![Go Report Card](https://travis-ci.org/SKatiyar/cri.svg?branch=develop)](https://travis-ci.org/SKatiyar/cri)

Package cri provides type-safe bindings for devtools protocol. It can be used with Chrome or any other target that implements the interface.


Protocol is generated by `cmd/generate.sh`. Script fetches latest version of protocol and generates types and domain (accessibility, domdebugger, performance etc.) packages. Master branch reflects tip of tree.

Tested with go1.4 and above.



### Install

```
go get -u github.com/SKatiyar/cri
```

### Usage

Taking a screenshot.

```go
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
  
	"github.com/SKatiyar/cri"
	"github.com/SKatiyar/cri/browser"
	"github.com/SKatiyar/cri/page"
)

func main() {
	conn, connErr := cri.NewConnection()
	if connErr != nil {
		fmt.Println(connErr)
		return
	}

	res, resErr := browser.New(conn).GetVersion()
	if resErr != nil {
		fmt.Println(resErr)
		return
	}

	pi := page.New(conn)
	if enableErr := pi.Enable(); enableErr != nil {
		fmt.Println(enableErr)
		return
	}

	nav, navErr := pi.Navigate(&page.NavigateRequest{
		Url: "https://www.example.com",
	})
	if navErr != nil {
		fmt.Println(navErr)
		return
	}

	pic, picErr := pi.CaptureScreenshot(nil)
	if picErr != nil {
		fmt.Println(picErr)
		return
	}

	img, imgErr := base64.StdEncoding.DecodeString(pic.Data)
	if imgErr != nil {
		fmt.Println(imgErr)
		return
	}

	if writeErr := ioutil.WriteFile("img.png", img, 0700); writeErr != nil {
		fmt.Println(writeErr)
		return
	}

	fmt.Println(res.JsVersion, nav.FrameId)
}
```

___

### TODO

* Add go get support to version 1.2 of protocol.
* Add tests for `connection.go`
* Add tests for domain packages.
* Simplify `On` function.
* Add timeout to `On` function.

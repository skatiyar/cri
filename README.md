# cri
Chrome Remote Interface for golang

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

	if writeErr := ioutil.WriteFile("img.jpg", img, 0700); writeErr != nil {
		fmt.Println(writeErr)
		return
	}

	fmt.Println(res.JsVersion, nav.FrameId)
}
```

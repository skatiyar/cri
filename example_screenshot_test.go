package cri_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/SKatiyar/cri"
	"github.com/SKatiyar/cri/emulation"
	"github.com/SKatiyar/cri/page"
)

// Example shows steps to take screenshot of a page.
func Example() {
	conn, connErr := cri.NewConnection(cri.SetEventTimeout(20 * time.Second))
	if connErr != nil {
		panic(connErr)
	}

	pi := page.New(conn)
	ei := emulation.New(conn)

	if enableErr := pi.Enable(); enableErr != nil {
		panic(enableErr)
	}
	if overideErr := ei.SetDeviceMetricsOverride(&emulation.SetDeviceMetricsOverrideRequest{
		Width:  1280,
		Height: 800,
	}); overideErr != nil {
		panic(overideErr)
	}

	urls := []string{
		"https://www.google.com",
		"https://www.chromestatus.com",
		"https://www.facebook.com",
		"https://www.example.com",
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(urls[i])
		_, navErr := pi.Navigate(&page.NavigateRequest{
			Url: urls[i],
		})
		if navErr != nil {
			panic(navErr)
		}

		_, loadErr := pi.LoadEventFired()
		if loadErr != nil {
			panic(loadErr.Error() + urls[i])
		}

		pic, picErr := pi.CaptureScreenshot(nil)
		if picErr != nil {
			panic(picErr)
		}

		img, imgErr := base64.StdEncoding.DecodeString(pic.Data)
		if imgErr != nil {
			panic(imgErr)
		}

		fileName := strconv.Itoa(i) + ".png"
		if writeErr := ioutil.WriteFile(fileName, img, 0700); writeErr != nil {
			panic(writeErr)
		}

		fmt.Println(fileName)

		if removeErr := os.Remove(fileName); removeErr != nil {
			panic(removeErr)
		}
	}

	// Unordered output:
	// https://www.google.com
	// 0.png
	// https://www.chromestatus.com
	// 1.png
	// https://www.facebook.com
	// 2.png
	// https://www.example.com
	// 3.png
}

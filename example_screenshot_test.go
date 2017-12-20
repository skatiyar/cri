package cri_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
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
		"www.google.com",
		"www.chromestatus.com",
		"www.facebook.com",
		"www.example.com",
	}

	for i := 0; i < len(urls); i++ {
		eventChn := make(chan page.LoadEventFiredParams)
		go func() {
			for {
				eve, loadErr := pi.LoadEventFired()
				if loadErr != nil {
					panic(loadErr.Error() + urls[i])
				}

				eventChn <- eve
				break
			}
		}()

		_, navErr := pi.Navigate(&page.NavigateRequest{
			Url: "https://" + urls[i],
		})
		if navErr != nil {
			panic(navErr)
		}

		<-eventChn

		pic, picErr := pi.CaptureScreenshot(nil)
		if picErr != nil {
			panic(picErr)
		}

		img, imgErr := base64.StdEncoding.DecodeString(pic.Data)
		if imgErr != nil {
			panic(imgErr)
		}

		fileName := urls[i] + ".png"
		if writeErr := ioutil.WriteFile(fileName, img, 0700); writeErr != nil {
			panic(writeErr)
		}

		fmt.Println(fileName)

		if removeErr := os.Remove(fileName); removeErr != nil {
			panic(removeErr)
		}
	}

	// Unordered output:
	// www.google.com.png
	// www.chromestatus.com.png
	// www.facebook.com.png
	// www.example.com.png
}

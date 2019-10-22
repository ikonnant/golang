// Run some code on play.golang.org and display the result
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var code = "golang"

// Errors are ignored for brevity.

func main() {
	// FireFox driver without specific version
	caps := selenium.Capabilities{"browserName": "firefox"}

	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()

	// Get simple playground interface
	_ = wd.Get("https://google.ru/")

	// Enter code in textarea
	elem, _ := wd.FindElement(selenium.ByCSSSelector, ".gLFyf.gsfi")
	_ = elem.Clear()
	_ = elem.SendKeys(code)

	time.Sleep(time.Millisecond * 10)

	// Click the run button
	btn, _ := wd.FindElement(selenium.ByCSSSelector, ".FPdoLc.VlcLAe .gNO89b")
	_ = btn.Click()

	// Get the result
	screen, _ := wd.Screenshot()

	file, err := os.Create("test.png")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	_, _ = file.Write(screen)
}

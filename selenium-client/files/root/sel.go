package main

import (
	"os"
	"log"
	"fmt"
	"time"

	"bitbucket.org/tebeka/selenium"
)

var code string = `
package main
import "fmt"

func main() {
	fmt.Println("Hello WebDriver!\n")
}
`

func main() {
	// Chrome driver without specific version
	wd, err := selenium.NewRemote(selenium.Capabilities{"browserName": "chrome"}, "http://"+os.Args[1]+":"+os.Args[2]+"/wd/hub")
	if err != nil {
		log.Fatalln(err)
	}
	defer wd.Quit()

	// Get simple playground interface
	wd.Get("http://play.golang.org/?simple=1")

	// Enter code in textarea
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		log.Fatalln(err)
	}
	elem.Clear()
	elem.SendKeys(code)

	// Click the run button
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		log.Fatalln(err)
	}
	btn.Click()

	// Get the result
	div, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		log.Fatalln(err)
	}

	output := ""
	// Wait for run to finish
	for {
		output, _ = div.Text()
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("Got: %s\n", output)
}


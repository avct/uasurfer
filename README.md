# User Agent Surfer

User Agent Surfer is a Go package that parses HTTP User-Agent strings with particular attention to the type of device, major OSes, and major browsers.

# Data

TBW

# Example

```
package main

import (
	"fmt"
	"github.com/avct/user-agent-surfer"
)

func main() {

	// Define a user agent string
	myUA := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36"

	// Instantiate new user_agent_surfer
	myBrowser := user_agent_surfer.New(myUA)

	// Print out some basic information
	fmt.Println("ua: ", myBrowser.UA)
	fmt.Println("browser: ", myBrowser.Browser.Name, " ", myBrowser.Browser.Version)
	fmt.Println("platform: ", myBrowser.Platform)
	fmt.Println("os: ", myBrowser.OS.Name, " ", myBrowser.OS.Version)
	fmt.Println("device type: ", myBrowser.DeviceType)

	// Simple logic example to determine Microsoft Surface RT (Tablet + Windows + modern-ish version of Windows)
	if myBrowser.DeviceType == "tablet" && myBrowser.Platform == "windows" && myBrowser.OS.Version > 6 {
		fmt.Println("The User-String string is probably Microsoft Surface RT.")
	} else {
		fmt.Println("The User-Agent string is definitely not Microsoft Surface RT.")
	}
}
```
# User Agent Surfer

User Agent Surfer is a Go package that parses and abstracts HTTP User-Agent strings with particular attention to speed, resource efficiency, and accuracy. Layout engine, browser language, and esoteric attributes are not parsed but are available in the BrowserProfile.UA string.

Approximately 98.5% of all web browsers used worldwide are identified.

# Implementation Example

```golang
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
		fmt.Println("The User-Agent string is probably Microsoft Surface RT.")
	} else {
		fmt.Println("The User-Agent string is definitely not Microsoft Surface RT.")
	}
}
```

# Browser Profile

The BrowserProfile supplies very specific enum string data along with integers for versions.

### Browser.Name
* `chrome`
* `safari`
* `ie`
* `firefox` (includes icecat, iceweasel, seamonkey)
* `android` - only Android 4.3 and earlier, 4.4 and later is `chrome`
* `opera`
* `silk`

### Browser.Version

tbw

### Platform
* `windows`
* `mac`
* `linux`
* `ipad`
* `iphone`
* `blackberry`
* `windows phone`
* `playstation`, `xbox`, `nintendo`

### OS.Name
* `2000`, `xp`, `vista`, `7`, `8`, `10`
* `os x`
* `ios`
* `playstation`, `xbox`, `nintendo`
* `chromeos`
* `webos`

### OS.Version

tbw

### DeviceType
* `computer`
* `phone`
* `tablet`
* `tv`
* `console`
* `wearable`

### Combination examples
Surface RT -> `windows`, `tablet`, OS.Version >= `6`
Android Tablet -> `android`, `tablet`
Microsoft Edge -> `ie`, Browser.Version == `12`

# To do

* Support [UC Browser](https://en.wikipedia.org/wiki/UC_Browser)
* Support NetFront
* Support Nokia
* Identify bots
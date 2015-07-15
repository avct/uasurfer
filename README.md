# User Agent Surfer

User Agent Surfer is a Go package that will parse and abstract HTTP User-Agent strings with particular attention to speed, resource efficiency, and accuracy. Layout engine, browser language, and esoteric attributes are not parsed but are available in the BrowserProfile.UA string.

Web browsers and operating systems that account for 98.5% of all worldwide use are identified.

# Example

```go
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

The BrowserProfile supplies specific enum strings along with integers for versions. The following strings should be supported, with the exception of linux OS being a mostly-hit but sometimes miss attribute.

#### Browser.Name
* `chrome`
* `safari`
* `ie`
* `firefox` (includes icecat, iceweasel, seamonkey)
* `android` - only Android ~4.3 and earlier use this name for the native WebView browser, 4.4 and later is `chrome`
* `opera`
* `ucbrowser`
* `silk`

#### Browser.Version

Browser.Version returns an integer of the correct top-level version attribute of the User-Agent String. The intention is to support math operators when evaluating versions. For example Chrome 45.0.23423 would return `45`.

#### Platform
* `windows`
* `mac`
* `linux` - Android OS uses Linux platform
* `ipad`
* `iphone`
* `blackberry`
* `windows phone`
* `kindle` - Fire models
* `playstation`, `xbox`, `nintendo`

#### OS.Name
* `2000`, `xp`, `vista`, `7`, `8`, `10`
* `os x`
* `ios`
* `android`
* `chromeos`
* `webos`
* `linux`
* `playstation`, `xbox`, `nintendo`

#### OS.Version

OS.Version returns an integer for the OS version, which is the NT major version for Windows (e.g. NT 6.2 is `6`) and minor version for OS X (e.g. OS X 10.11.6 is `11`). This is to allow ease of use around math operators the version numbers. Here are some examples across the platform, os.name, and os.version:

* For Windows XP (Windows NT 5.1), "`windows`" is the platform, "`xp`" is the name, and `5` the version.
* For OS X 10.5.1, "`mac`" is the platform, "`os x`" the name, and `5` the version.
* For Android 5.1, "`linux`" is the platform, "`android`" is the name, and `5` the version.
* For iOS 5.1, "`iphone`" or "`ipad`" is the platform, "`ios`" is the name, and `5` the version.

#### DeviceType
DeviceType is typically quite accurate, though determining between phones and tablets on Android is not always possible due to how some vendors design their UA strings. A mobile Android device without tablet indication is classified a phone.

* `computer`
* `phone`
* `tablet`
* `tv`
* `console`
* `wearable`

#### Combination examples
* Surface RT -> `windows`, `tablet`, OS.Version >= `6`
* Android Tablet -> `android`, `tablet`
* Microsoft Edge -> `ie`, Browser.Version == `12`

# To do

* Add tests for UC Browser
* Support NetFront
* Support [Nokia Browser](https://en.wikipedia.org/wiki/Nokia_Browser_for_Symbian)
* Identify bots
* Add OS->Browser identification logic
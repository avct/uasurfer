# User Agent Surfer

User Agent Surfer (uasurfer) is a Go package that will parse and abstract HTTP User-Agent strings with particular attention to accuracy, speed, and resource efficiency. The following information is returned by uasurfer after supplying it a raw UA string:

* **Browser name** (e.g. Chrome)
* **Browser major version** (e.g. Chrome 45)
* **Platform** (e.g. iPad)
* **OS name** (e.g. iOS)
* **OS major version** (e.g. iOS 9)
* **Device type** (e.g. tablet)

Layout engine, browser language, and other esoteric attributes are not parsed.

Web browsers and operating systems that account for 98.5% of all worldwide use are identified.

# Browser Profile

The BrowserProfile supplies specific const (int)  along with integers for versions. The following strings should be supported, with the exception of linux OS being a mostly-hit but sometimes miss attribute.

#### Browser.Name
* `chrome`
* `safari`
* `ie`
* `firefox` (includes icecat, iceweasel, seamonkey)
* `android` - only Android ~4.3 and earlier can use this name for the native WebView browser, 4.4 and later is always `chrome`
* `opera`
* `ucbrowser`
* `silk`
* `nokia`
* `gsa` - Google search app on iOS
* `spotify` - applicable for advertising applications

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
* `bot`

#### OS.Name
* `2000`, `xp`, `vista`, `7`, `8`, `10`
* `os x`
* `ios`
* `android`
* `chromeos`
* `webos`
* `linux`
* `playstation`, `xbox`, `nintendo`
* `bot`

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
* `bot`

#### Combination examples
* Surface RT -> `windows`, `tablet`, OS.Version >= `6`
* Android Tablet -> `android`, `tablet`
* Microsoft Edge -> `ie`, Browser.Version == `12`

# To do

* Support bots
* Support NetFront
* Support Nokia browser
* Support Kindle Fire
* Add OS->Browser identification logic
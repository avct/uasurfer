# User Agent Surfer

User Agent Surfer (uasurfer) is a Go package that parses and abstract HTTP User-Agent strings with particular attention to accuracy, speed, and resource efficiency. The following information is returned by uasurfer after supplying it a raw UA string:

* **Browser name** (e.g. `chrome`)
* **Browser major version** (e.g. `45`)
* **Platform** (e.g. `ipad`)
* **OS name** (e.g. `ios`)
* **OS major version** (e.g. `9`)
* **Device type** (e.g. `tablet`)

Layout engine, browser language, and other esoteric attributes are not parsed.

Web browsers and operating systems that account for 98.5% of all worldwide use are identified.

## Parse(ua string) Function

The `uasurfer.Parse()` function accepts a user agent string (string) and returns specific named constants along with integers (unint8) for versions.

```
// Define a user agent string
myUA := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36"

// Parse() is multivariate, including returning the full UA string last
browserName, browserVersion, platform, osName, osVersion, deviceType, ua := uasurfer.Parse(myUA)
```

**Usage note:** There are some minor OSes that do no return a version, see docs below, and linux OS can be hit-or-miss at this stage given the plethura of OS names. Linux as a platform is quite accurate.

#### Browser Name
* Google Chrome, Chromium - `BrowserChrome`
* Apple Safari, Google Search Application (GSA) - `BrowserSafari`
* Microsoft Internet Explorer & Edge - `BrowserIE`
* Mozilla Firefox, Icecat, Iceweasel, Seamonkey - `BrowserFirefox`
* Android WebView (Android <4.4) - `BrowserAndroid`
* Opera - `BrowserOpera`
* UCBrowser - `BrowserUCBrowser`
* Amazon Silk - `BrowserSilk`
* Spotify - `BrowserSpotify`
* RIM Blackberry - `BrowserBlackberry`
* Unknown - `BrowserUnknown`

#### Browser Version

Browser version returns an `unint8` of the correct top-level version attribute of the User-Agent String. For example Chrome 45.0.23423 would return `45`. The intention is to support math operators with versions, such as Chrome version >23.

Unknown version is returned as `0`.

#### Platform
* Microsoft Windows - `PlatformWindows`
* Apple Macintosh - `PlatformMac`
* Linux - `PlatformLinux` - Android OS uses Linux platform
* Apple iPad - `PlatformiPad`
* Apple iPhone - `PlatformiPhone`
* RIM Blackberry - `PlatformBlackberry`
* Microsoft Windows Phone - `PlatformWindowsPhone`
* Amazon Kindle & Kindle Fire - `PlatformKindle` - Fire models
* Sony Playstation, Vita, PSP - `PlatformPlaystation`
* Microsoft Xbox - `PlatformXbox`
* Nintendo DS, Wii, etc. - `PlatformNintendo`
* Unknown - `PlatformUnknown`

#### OS Name
* `2000`, `xp`, `vista`, `7`, `8`, `10`
* `os x`
* `ios`
* `android`
* `chromeos`
* `webos`
* `linux`
* `playstation`, `xbox`, `nintendo`
* `unknown`

#### OS Version

OS version will be an integer (unint8) for the mjor OS version, which is the NT major version for Windows (e.g. NT 6.2 is `6`) and minor version for OS X (e.g. OS X 10.11.6 is `11`). `0` indicates the OS verison is unknown, or not evaluated. This is to allow ease of use around math operators the version numbers. Here are some examples across the platform, os.name, and os.version:

* For Windows XP (Windows NT 5.1), "`PlatformWindows`" is the platform, "`OSWindowsXP`" is the name, and `5` the version.
* For OS X 10.5.1, "`PlatformMac`" is the platform, "`OSMacOSX`" the name, and `5` the version.
* For Android 5.1, "`PlatformLinux`" is the platform, "`OSAndroid`" is the name, and `5` the version.
* For iOS 5.1, "`PlatformiPhone`" or "`PlatformiPad`" is the platform, "`OSiOS`" is the name, and `5` the version.

#### DeviceType
DeviceType is typically quite accurate, though determining between phones and tablets on Android is not always possible due to how some vendors design their UA strings. A mobile Android device without tablet indicator defaults to being classified as a phone. DeviceTV supports major brands like Philips, Sharp, Vizio and steaming boxes such as Apple, Google, Roku, Amazon.

* `DeviceComputer`
* `DevicePhone`
* `DeviceTablet`
* `DeviceTV`
* `DeviceConsole`
* `DeviceWearable`
* `DeviceUnknown`

## Example Combinations of Attributes
* Surface RT -> `windows`, `tablet`, OS.Version >= `6`
* Android Tablet -> `android`, `tablet`
* Microsoft Edge -> `ie`, Browser.Version == `12`

## To do

* Support version on Firefox derivatives (e.g. SeaMonkey)
* Support bots
* Support NetFront
* Support Nokia browser
* iOS safari browser identification based on iOS version
* Add android version to browser identification
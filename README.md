# User Agent Surfer

User Agent Surfer (uasurfer) is a Go package that parses and abstract HTTP User-Agent strings with particular attention to accuracy, speed, and resource efficiency. Each full parse requires `608 B/op` & `11 allocs/op`. The following information is returned by uasurfer after supplying it a raw UA string:

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
* `BrowserChrome` - Google [Chrome](https://en.wikipedia.org/wiki/Google_Chrome), [Chromium](https://en.wikipedia.org/wiki/Chromium_(web_browser))
* `BrowserSafari` - Apple [Safari](https://en.wikipedia.org/wiki/Safari_(web_browser)), Google Search ([GSA](https://itunes.apple.com/us/app/google/id284815942))
* `BrowserIE` - Microsoft [Internet Explorer](https://en.wikipedia.org/wiki/Internet_Explorer), [Edge](https://en.wikipedia.org/wiki/Microsoft_Edge)
* `BrowserFirefox` - Mozilla [Firefox](https://en.wikipedia.org/wiki/Firefox), GNU [IceCat](https://en.wikipedia.org/wiki/GNU_IceCat), [Iceweasel](https://en.wikipedia.org/wiki/Mozilla_Corporation_software_rebranded_by_the_Debian_project#Iceweasel), [Seamonkey](https://en.wikipedia.org/wiki/SeaMonkey)
* `BrowserAndroid` - Android [WebView](https://developer.chrome.com/multidevice/webview/overview) (Android OS <4.4 only)
* `BrowserOpera` - [Opera](https://en.wikipedia.org/wiki/Opera_(web_browser))
* `BrowserUCBrowser` - [UC Browser](https://en.wikipedia.org/wiki/UC_Browser)
* `BrowserSilk` - Amazon [Silk](https://en.wikipedia.org/wiki/Amazon_Silk)
* `BrowserSpotify` - [Spotify](https://en.wikipedia.org/wiki/Spotify#Clients) desktop client
* `BrowserBlackberry` - RIM [BlackBerry](https://en.wikipedia.org/wiki/BlackBerry)
* `BrowserUnknown` - Unknown

#### Browser Version

Browser version returns an `unint8` of the major version attribute of the User-Agent String. For example Chrome 45.0.23423 would return `45`. The intention is to support math operators with versions, such as "do XYZ for Chrome version >23".

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
* `OSWindows2000`, `OSWindowsXP`, `OSWindowsVista`, `OSWindows7`, `OSWindows8`, `OSWindows10`
* `OSMacOSX`
* `OSiOS`
* `OSAndroid`
* `OSChromeOS`
* `OSWebOS`
* `OSLinux`
* `OSPlaystation`
* `OSXbox`
* `OSNintendo`
* `OSUnknown`

#### OS Version

OS version will be an integer (unint8) for the mjor OS version, which is the NT major version for Windows (e.g. NT 6.2 is `6`) and minor version for OS X (e.g. OS X 10.11.6 is `11`). `0` indicates the OS verison is unknown, or not evaluated. This is to allow ease of use around math operators the version numbers. Here are some examples across the platform, os.name, and os.version:

* For Windows XP (Windows NT 5.1), "`PlatformWindows`" is the platform, "`OSWindowsXP`" is the name, and `5` the version.
* For OS X 10.5.1, "`PlatformMac`" is the platform, "`OSMacOSX`" the name, and `5` the version.
* For Android 5.1, "`PlatformLinux`" is the platform, "`OSAndroid`" is the name, and `5` the version.
* For iOS 5.1, "`PlatformiPhone`" or "`PlatformiPad`" is the platform, "`OSiOS`" is the name, and `5` the version.

###### Windows Version Guide

OSWindows versioning uses the "NT" version of an OS as listed in the User Agent String. These versions of Windows are supported and return the associated `unint8`:

* Windows 10 - `10`
* Windows 8, 8.1 - `8`
* Windows 7 - `7`
* Windows Vista - `6`
* Windows XP - `5`
* Windows 2000 - `4`

Windows 95, 98, and ME represent 0.01% of traffic worldwide and are not available through this package.

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
* Surface RT -> `OSWindows8`, `DeviceTablet`, OSVersion >= `6`
* Android Tablet -> `OSAndroid`, `DeviceTablet`
* Microsoft Edge -> `BrowserIE`, BrowserVersion == `12`

## To do

* Remove compiled regexp in favor of string.Contains wherever possible (lowers mem/alloc)
* Support version on Firefox derivatives (e.g. SeaMonkey)
* Support bots
* Potential additional browser support:
 * "NetFront" (1% share in India)
 * "QQ Browser" (6.5% share in China)
 * "Sogou Explorer" (5% share in China)
 * "Maxthon" (1.5% share in China)
 * "Nokia"
* Potential additional OS support:
 * "Nokia" (5% share in India)
 * "Series 40" (5.5% share in India)
* iOS safari browser identification based on iOS version
* Add android version to browser identification
package uasurfer

import (
	// "bufio"
	// "fmt"
	// "os"
	"testing"
)

// Test values are ordered: BrowserName, BrowserVersion, OSName, OSVersion, DeviceType
var testUAVars = []struct {
	UA             string
	browserName    BrowserName
	browserVersion int
	Platform       Platform
	osName         OSName
	osVersion      int
	deviceType     DeviceType
}{
	// Empty
	{"",
		BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DeviceUnknown},

	// Single char
	{"a",
		BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DeviceUnknown},

	// Some random string
	{"some random string",
		BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DeviceUnknown},

	// Potentially malformed ua
	{")(",
		BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DeviceUnknown},

	// iPhone
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		BrowserSafari, 6, PlatformiPhone, OSiOS, 7, DevicePhone},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		BrowserSafari, 8, PlatformiPhone, OSiOS, 8, DevicePhone},

	// iPad
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		BrowserSafari, 4, PlatformiPad, OSiOS, 3, DeviceTablet},

	{"Mozilla/5.0 (iPad; CPU OS 9_0 like Mac OS X) AppleWebKit/601.1.17 (KHTML, like Gecko) Version/8.0 Mobile/13A175 Safari/600.1.4",
		BrowserSafari, 8, PlatformiPad, OSiOS, 9, DeviceTablet},

	// Chrome
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36",
		BrowserChrome, 43, PlatformMac, OSMacOSX, 10, DeviceComputer},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/534.48.3",
		BrowserChrome, 19, PlatformiPhone, OSiOS, 5, DevicePhone},

	{"Mozilla/5.0 (Linux; Android 6.0; Nexus 5X Build/MDB08L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
		BrowserChrome, 46, PlatformLinux, OSAndroid, 6, DevicePhone},

	// Chromium (Chrome)
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Ubuntu/11.10 Chromium/18.0.1025.142 Chrome/18.0.1025.142 Safari/535.19",
		BrowserChrome, 18, PlatformLinux, OSLinux, 0, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		BrowserChrome, 45, PlatformMac, OSMacOSX, 11, DeviceComputer},

	//TODO: refactor "getMajorVersion()" to handle this device/chrome version douchebaggery
	// {"Mozilla/5.0 (Linux; Android 4.4.2; en-gb; SAMSUNG SM-G800F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.6 Chrome/28.0.1500.94 Mobile Safari/537.36",
	// 	BrowserChrome, 28, PlatformLinux, OSAndroid, 4, DevicePhone},

	// Safari
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		BrowserSafari, 8, PlatformMac, OSMacOSX, 10, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		BrowserSafari, 3, PlatformMac, OSMacOSX, 5, DeviceComputer},

	// Firefox
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
		BrowserFirefox, 1, PlatformiPhone, OSiOS, 8, DevicePhone},

	{"Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0",
		BrowserFirefox, 41, PlatformLinux, OSAndroid, 4, DeviceTablet},

	{"Mozilla/5.0 (Android; Mobile; rv:40.0) Gecko/40.0 Firefox/40.0",
		BrowserFirefox, 40, PlatformLinux, OSAndroid, 0, DevicePhone},

	{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:38.0) Gecko/20100101 Firefox/38.0",
		BrowserFirefox, 38, PlatformLinux, OSLinux, 0, DeviceComputer},

	// Silk
	{"Mozilla/5.0 (Linux; U; Android 4.4.3; de-de; KFTHWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.47 like Chrome/37.0.2026.117 Safari/537.36",
		BrowserSilk, 3, PlatformLinux, OSKindle, 4, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; en-us; KFJWI Build/IMM76D) AppleWebKit/535.19 (KHTML like Gecko) Silk/2.4 Safari/535.19 Silk-Acceleratedtrue",
		BrowserSilk, 2, PlatformLinux, OSKindle, 0, DeviceTablet},

	// Opera
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 OPR/18.0.1284.68",
		BrowserOpera, 18, PlatformWindows, OSWindows, 7, DeviceComputer},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) OPiOS/10.2.0.93022 Mobile/12H143 Safari/9537.53",
		BrowserOpera, 10, PlatformiPhone, OSiOS, 8, DevicePhone},

	// Internet Explorer -- https://msdn.microsoft.com/en-us/library/hh869301(v=vs.85).aspx
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.123",
		BrowserIE, 12, PlatformWindows, OSWindows, 10, DeviceComputer},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		BrowserIE, 10, PlatformWindows, OSWindows, 8, DeviceComputer},

	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		BrowserIE, 11, PlatformWindows, OSWindows, 8, DeviceComputer},

	{"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; DEVICE INFO) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.123",
		BrowserIE, 12, PlatformWindowsPhone, OSWindowsPhone, 10, DevicePhone},

	{"Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 520) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537",
		BrowserIE, 11, PlatformWindowsPhone, OSWindowsPhone, 8, DevicePhone},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		BrowserIE, 5, PlatformWindows, OSWindows, 4, DeviceComputer},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		BrowserIE, 7, PlatformWindows, OSWindows, 7, DeviceComputer},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)", //Windows Surface RT tablet
		BrowserIE, 10, PlatformWindows, OSWindows, 8, DeviceTablet},

	// UC Browser
	{"Mozilla/5.0 (Linux; U; Android 2.3.4; en-US; MT11i Build/4.0.2.A.0.62) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.0.1.275 U3/0.8.0 Mobile Safari/534.31",
		BrowserUCBrowser, 9, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.0.4; en-US; Micromax P255 Build/IMM76D) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.2.0.308 U3/0.8.0 Mobile Safari/534.31",
		BrowserUCBrowser, 9, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"UCWEB/2.0 (Java; U; MIDP-2.0; en-US; MicromaxQ5) U2/1.0.0 UCBrowser/9.4.0.342 U2/1.0.0 Mobile",
		BrowserUCBrowser, 9, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// Nokia Browser
	// {"Mozilla/5.0 (Series40; Nokia501/14.0.4/java_runtime_version=Nokia_Asha_1_2; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/4.0.0.0.45",
	// 	BrowserUnknown, 4, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// {"Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaN8-00/111.040.1511; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/535.1 (KHTML, like Gecko) NokiaBrowser/8.3.1.4 Mobile Safari/535.1",
	// 	BrowserUnknown, 8, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// {"NokiaN97/21.1.107 (SymbianOS/9.4; Series60/5.0 Mozilla/5.0; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebkit/525 (KHTML, like Gecko) BrowserNG/7.1.4",
	// 	BrowserUnknown, 7, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// ChromeOS
	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		BrowserChrome, 4, PlatformLinux, OSChromeOS, 0, DeviceComputer},

	// WebOS
	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; de-DE) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
		BrowserUnknown, 0, PlatformLinux, OSWebOS, 0, DeviceTablet},

	{"Mozilla/5.0 (webOS/1.4.1.1; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		BrowserUnknown, 1, PlatformLinux, OSWebOS, 0, DevicePhone},

	// Android WebView (Android <= 4.3)
	{"Mozilla/5.0 (Linux; U; Android 2.2; en-us; DROID2 GLOBAL Build/S273) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.0.3; de-ch; HTC Sensation Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari53/4.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	// BlackBerry
	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+",
		BrowserBlackberry, 7, PlatformBlackberry, OSBlackberry, 0, DeviceTablet},

	{"Mozilla/5.0 (BB10; Kbd) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.2.1.1925 Mobile Safari/537.35+",
		BrowserBlackberry, 10, PlatformBlackberry, OSBlackberry, 0, DevicePhone},

	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) BlackBerry8703e/4.1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/104",
		BrowserBlackberry, 0, PlatformBlackberry, OSBlackberry, 0, DevicePhone},

	// Windows Phone
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 625; ANZ941)",
		BrowserIE, 10, PlatformWindowsPhone, OSWindowsPhone, 8, DevicePhone},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 900)",
		BrowserIE, 9, PlatformWindowsPhone, OSWindowsPhone, 7, DevicePhone},

	// Kindle eReader
	{"Mozilla/5.0 (Linux; U; en-US) AppleWebKit/528.5+ (KHTML, like Gecko, Safari/528.5+) Version/4.0 Kindle/3.0 (screen 600×800; rotate)",
		BrowserUnknown, 4, PlatformLinux, OSKindle, 0, DeviceTablet},

	{"Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+",
		BrowserUnknown, 5, PlatformLinux, OSKindle, 0, DeviceTablet},

	// Amazon Fire
	{"Mozilla/5.0 (Linux; U; Android 4.4.3; de-de; KFTHWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.67 like Chrome/39.0.2171.93 Safari/537.36",
		BrowserSilk, 3, PlatformLinux, OSKindle, 4, DeviceTablet}, // Fire tablet

	{"Mozilla/5.0 (Linux; U; Android 4.2.2; en­us; KFTHWI Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.22 like Chrome/34.0.1847.137 Mobile Safari/537.36",
		BrowserSilk, 3, PlatformLinux, OSKindle, 4, DeviceTablet}, // Fire tablet, but with "Mobile"

	{"Mozilla/5.0 (Linux; Android 4.4.4; SD4930UR Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/34.0.0.0 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/35.0.0.48.273;]",
		BrowserChrome, 34, PlatformLinux, OSKindle, 4, DevicePhone}, // Facebook app on Fire Phone

	// extra logic to identify phone when using silk has not been added
	// {"Mozilla/5.0 (Linux; Android 4.4.4; SD4930UR Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.67 like Chrome/39.0.2171.93 Mobile Safari/537.36",
	// 	BrowserSilk, 3, PlatformLinux, OSKindle, 4, DevicePhone}, // Silk on Fire Phone

	// Nintendo
	{"Opera/9.30 (Nintendo Wii; U; ; 2047-7; fr)",
		BrowserOpera, 9, PlatformNintendo, OSNintendo, 0, DeviceConsole},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US",
		BrowserUnknown, 0, PlatformNintendo, OSNintendo, 0, DeviceConsole},

	// Xbox
	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)", //Xbox 360
		BrowserIE, 9, PlatformXbox, OSXbox, 6, DeviceConsole},

	// Playstation

	{"Mozilla/5.0 (Playstation Vita 1.61) AppleWebKit/531.22.8 (KHTML, like Gecko) Silk/3.2",
		BrowserSilk, 3, PlatformPlaystation, OSPlaystation, 0, DeviceConsole},

	// Smart TVs and TV dongles
	{"Mozilla/5.0 (CrKey armv7l 1.4.15250) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.0 Safari/537.36", // Chromecast
		BrowserChrome, 31, PlatformUnknown, OSUnknown, 0, DeviceTV},

	{"Mozilla/5.0 (Linux; GoogleTV 3.2; VAP430 Build/MASTER) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.77 Safari/534.24", // Google TV
		BrowserChrome, 11, PlatformLinux, OSUnknown, 0, DeviceTV},

	{"Mozilla/5.0 (Linux; Android 5.0; ADT-1 Build/LPX13D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Mobile Safari/537.36", // Android TV
		BrowserChrome, 40, PlatformLinux, OSAndroid, 5, DeviceTV},

	{"Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22", // Amazon Fire
		BrowserChrome, 25, PlatformLinux, OSAndroid, 4, DeviceTV},

	{"Mozilla/5.0 (Unknown; Linux armv7l) AppleWebKit/537.1+ (KHTML, like Gecko) Safari/537.1+ LG Browser/6.00.00(+mouse+3D+SCREEN+TUNER; LGE; GLOBAL-PLAT5; 03.07.01; 0x00000001;); LG NetCast.TV-2013/03.17.01 (LG, GLOBAL-PLAT4, wired)", // LG TV
		BrowserUnknown, 0, PlatformLinux, OSUnknown, 0, DeviceTV},

	{"Mozilla/5.0 (X11; FreeBSD; U; Viera; de-DE) AppleWebKit/537.11 (KHTML, like Gecko) Viera/3.10.0 Chrome/23.0.1271.97 Safari/537.11", // Panasonic Viera
		BrowserChrome, 23, PlatformLinux, OSLinux, 0, DeviceTV},

	{"Mozilla/5.0 (DTV) AppleWebKit/531.2+ (KHTML, like Gecko) Espial/6.1.5 AQUOSBrowser/2.0 (US01DTV;V;0001;0001)", // Sharp Aquos
		BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DeviceTV},

	{"Roku/DVP-5.2 (025.02E03197A)", // Roku
		BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DeviceTV},

	// Google search app (GSA) for iOS -- it's Safari in disguise as of v6
	{"Mozilla/5.0 (iPad; CPU OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) GSA/6.0.51363 Mobile/12F69 Safari/600.1.4",
		BrowserSafari, 8, PlatformiPad, OSiOS, 8, DeviceTablet},

	// Spotify (applicable for advertising applications)
	{"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Spotify/1.0.9.133 Safari/537.36",
		BrowserSpotify, 1, PlatformWindows, OSWindows, 5, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Spotify/1.0.9.133 Safari/537.36",
		BrowserSpotify, 1, PlatformMac, OSMacOSX, 10, DeviceComputer},

	// Bots
	// {"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	// 	"bot", 0, "bot", "bot", 0, DeviceBot},

	// {"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	// 	"bot", 0, "bot", "bot", 0, DeviceBot},

	// Unknown or partially handled
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3", //Seamonkey (~FF)
		BrowserFirefox, 0, PlatformMac, OSMacOSX, 4, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre", //Camino (~FF)
		BrowserUnknown, 0, PlatformMac, OSMacOSX, 5, DeviceComputer},

	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", //firefox OS
		BrowserFirefox, 26, PlatformUnknown, OSUnknown, 0, DevicePhone},

	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.45 Safari/535.19", //chrome for android having requested desktop site
		BrowserChrome, 18, PlatformLinux, OSLinux, 0, DeviceComputer},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		BrowserOpera, 10, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// BrowserQQ
	// {"Mozilla/5.0 (Windows NT 6.2; WOW64; Trident/7.0; Touch; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 2.0.50727; .NET CLR 3.0.30729; InfoPath.3; Tablet PC 2.0; QQBrowser/7.6.21433.400; rv:11.0) like Gecko",
	// 	BrowserQQ, 7, PlatformWindows, OSWindows, 8, DeviceTablet},

	// {"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36 QQBrowser/9.0.2191.400",
	// 	BrowserQQ, 9, PlatformWindows, OSWindows, 7, DeviceComputer},

	// SUBMITTED TESTS
	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; dream) AppleWebKit/525.10+ (KHTML,like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; generic) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.0.3; de-de; A80KSC Build/ECLAIR) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.5; en-gb; T-Mobile G1 Build/CRC1) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.5; es-; FBW1_4 Build/MASTER) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux U; Android 1.5 en-us hero) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.5; en-us; Opus One Build/RBE.00.00) AppleWebKit/528.18.1 (KHTML, like Gecko) Version/3.1.1 Mobile Safari/525.20.1",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.6; ar-us; SonyEricssonX10i Build/R2BA026) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	// TODO: support names of Android OS?
	{"Mozilla/5.0 (Linux; U; Android Donut; de-de; HTC Tattoo 1.52.161.1 Build/Donut) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.6; en-gb; HTC Tattoo Build/DRC79) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 1.6; ja-jp; Docomo HT-03A Build/DRD08) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 1, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; Nexus One Build/ERD62) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 2.1-update1; en-au; HTC_Desire_A8183 V1.16.841.1 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; generic) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 2, DevicePhone},

	// TODO support named versions of Android?
	{"Mozilla/5.0 (Linux; U; Android Eclair; en-us; sholes) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		BrowserAndroid, 3, PlatformLinux, OSAndroid, 0, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 2.2; en-sa; HTC_DesireHD_A9191 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 2.2.1; en-gb; HTC_DesireZ_A7272 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	// INCORRECT DEVICE --> PHONE
	{"Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; Sensation_4G Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Safari/533.16",
		BrowserAndroid, 5, PlatformLinux, OSAndroid, 2, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; Android 2.3.5; ko-kr; SHW-M250S Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 2.3.7; ja-jp; L-02D Build/GWK74) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 3.0; xx-xx; Transformer TF101 Build/HRI66) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 3, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 3, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; Android 3.2; de-de; A100 Build/HTJ85B) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 3, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; Android 4.0.1; en-us; sdk Build/ICS_MR0) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	// TODO support "android-" version prefix
	// {"Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7",
	// 	BrowserChrome, 16, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.0.3; pl-pl; Transformer TF101 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Nexus S Build/JRO03E) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.1; en-gb; Build/JRN84D) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Mobile Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; el-gr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.2; en-us; Nexus 10 Build/JVP15I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; Android 4.2; ro-ro; LT18i Build/4.1.B.0.431) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		BrowserAndroid, 4, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JWR66D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.111 Safari/537.36",
		BrowserChrome, 27, PlatformLinux, OSAndroid, 4, DeviceTablet},

	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 7 Build/KOT24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Safari/537.36",
		BrowserChrome, 30, PlatformLinux, OSAndroid, 4, DeviceTablet},

	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 4 Build/KRT16E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Mobile Safari",
		BrowserChrome, 30, PlatformLinux, OSAndroid, 4, DevicePhone},

	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) BlackBerry8703e/4.1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/104",
		BrowserBlackberry, 0, PlatformBlackberry, OSBlackberry, 0, DevicePhone},

	{"Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.1.0.4633 Mobile Safari/537.10+",
		BrowserBlackberry, 10, PlatformBlackberry, OSBlackberry, 0, DevicePhone},

	{"Mozilla/5.0 (BB10; Kbd) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.2.1.1925 Mobile Safari/537.35+",
		BrowserBlackberry, 10, PlatformBlackberry, OSBlackberry, 0, DevicePhone},

	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 1.0.0; en-US) AppleWebKit/534.11 (KHTML, like Gecko) Version/7.1.0.7 Safari/534.11",
		BrowserBlackberry, 7, PlatformBlackberry, OSBlackberry, 0, DeviceTablet},

	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+",
		BrowserBlackberry, 7, PlatformBlackberry, OSBlackberry, 0, DeviceTablet},

	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		BrowserChrome, 4, PlatformLinux, OSChromeOS, 0, DeviceComputer},

	{"Mozilla/5.0 (X11; CrOS armv7l 5500.100.6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.120 Safari/537.36",
		BrowserChrome, 34, PlatformLinux, OSChromeOS, 0, DeviceComputer},

	// {"Mozilla/5.0 (Mobile; rv:14.0) Gecko/14.0 Firefox/14.0",
	// 	BrowserFirefox, 14, OSFirefoxOS, 14, DevicePhone},

	// {"Mozilla/5.0 (Mobile; rv:17.0) Gecko/17.0 Firefox/17.0",
	// 	BrowserFirefox, , OSFirefoxOS, DevicePhone},

	// {"Mozilla/5.0 (Mobile; rv:18.1) Gecko/18.1 Firefox/18.1",
	// 	BrowserFirefox, , OSFirefoxOS, DevicePhone},

	// {"Mozilla/5.0 (Tablet; rv:18.1) Gecko/18.1 Firefox/18.1",
	// 	BrowserFirefox, , OSFirefoxOS, DevicePhone},

	// {"Mozilla/5.0 (Mobile; LG-D300; rv:18.1) Gecko/18.1 Firefox/18.1",
	// 	BrowserFirefox, , OSFirefoxOS, DevicePhone},

	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		BrowserSafari, 4, PlatformiPad, OSiOS, 3, DeviceTablet},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
		BrowserSafari, 4, PlatformiPhone, OSiOS, 4, DevicePhone},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		BrowserSafari, 5, PlatformiPhone, OSiOS, 5, DevicePhone},

	{"Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		BrowserSafari, 5, PlatformiPad, OSiOS, 5, DeviceTablet},

	{"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		BrowserSafari, 6, PlatformiPad, OSiOS, 6, DeviceTablet},

	// TODO handle default browser based on iOS version
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A5376e",
		BrowserUnknown, 0, PlatformiPhone, OSiOS, 6, DevicePhone},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		BrowserSafari, 6, PlatformiPhone, OSiOS, 7, DevicePhone},

	{"Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		BrowserSafari, 7, PlatformiPad, OSiOS, 7, DeviceTablet},

	{"Mozilla/5.0 (iPad; CPU OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
		BrowserSafari, 7, PlatformiPad, OSiOS, 7, DeviceTablet},

	// TODO handle default browser based on iOS version
	// {"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/538.34.9 (KHTML, like Gecko) Mobile/12A4265u",
	// 	BrowserSafari, 8, PlatformiPhone, OSiOS, 8, DevicePhone},

	// TODO extrapolate browser from iOS version
	// {"Mozilla/5.0 (iPad; CPU OS 8_0 like Mac OS X) AppleWebKit/538.34.9 (KHTML, like Gecko) Mobile/12A4265u",
	// 	BrowserSafari, 8, PlatformiPad, OSiOS, 8, DeviceTablet},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		BrowserSafari, 8, PlatformiPhone, OSiOS, 8, DevicePhone},

	{"Mozilla/5.0 (X11; U; Linux x86_64; en; rv:1.9.0.14) Gecko/20080528 Ubuntu/9.10 (karmic) Epiphany/2.22 Firefox/3.0",
		BrowserFirefox, 3, PlatformLinux, OSLinux, 0, DeviceComputer},

	// Can't parse browser due to limitation of user agent library
	{"Mozilla/5.0 (X11; U; Linux x86_64; zh-TW; rv:1.9.0.8) Gecko/2009032712 Ubuntu/8.04 (hardy) Firefox/3.0.8 GTB5",
		BrowserFirefox, 3, PlatformLinux, OSLinux, 0, DeviceComputer},

	{"Mozilla/5.0 (compatible; Konqueror/3.5; Linux; x86_64) KHTML/3.5.5 (like Gecko) (Debian)",
		BrowserUnknown, 0, PlatformLinux, OSLinux, 0, DeviceComputer},

	{"Mozilla/5.0 (X11; U; Linux i686; de; rv:1.9.1.5) Gecko/20091112 Iceweasel/3.5.5 (like Firefox/3.5.5; Debian-3.5.5-1)",
		BrowserFirefox, 3, PlatformLinux, OSLinux, 0, DeviceComputer},

	// TODO consider bot?
	// {"Miro/2.0.4 (http://www.getmiro.com/; Darwin 10.3.0 i386)",
	// 	BrowserUnknown, 0, PlatformMac, OSMacOSX, 3, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3",
		BrowserFirefox, 0, PlatformMac, OSMacOSX, 4, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		BrowserSafari, 3, PlatformMac, OSMacOSX, 5, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre",
		BrowserUnknown, 0, PlatformMac, OSMacOSX, 5, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_2; en-US) AppleWebKit/533.1 (KHTML, like Gecko) Chrome/5.0.329.0 Safari/533.1",
		BrowserChrome, 5, PlatformMac, OSMacOSX, 6, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 (.NET CLR 3.5.30729)",
		BrowserFirefox, 3, PlatformMac, OSMacOSX, 6, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7",
		BrowserSafari, 5, PlatformMac, OSMacOSX, 7, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:9.0) Gecko/20111222 Thunderbird/9.0.1",
		BrowserUnknown, 0, PlatformMac, OSMacOSX, 7, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		BrowserChrome, 16, PlatformMac, OSMacOSX, 7, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/535.18.5 (KHTML, like Gecko) Version/5.2 Safari/535.18.5",
		BrowserSafari, 5, PlatformMac, OSMacOSX, 8, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
		BrowserChrome, 4, PlatformMac, OSMacOSX, 8, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9) AppleWebKit/537.35.1 (KHTML, like Gecko) Version/6.1 Safari/537.35.1",
		BrowserSafari, 6, PlatformMac, OSMacOSX, 9, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.34.48 (KHTML, like Gecko) Version/8.0 Safari/538.35.8",
		BrowserSafari, 8, PlatformMac, OSMacOSX, 10, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.32 (KHTML, like Gecko) Version/7.1 Safari/538.4",
		BrowserSafari, 7, PlatformMac, OSMacOSX, 10, DeviceComputer},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		BrowserOpera, 10, PlatformUnknown, OSUnknown, 0, DevicePhone},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		BrowserOpera, 10, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// TODO: support OneBrowser? https://play.google.com/store/apps/details?id=com.tencent.ibibo.mtt&hl=en_GB
	// {"OneBrowser/3.1 (NokiaN70-1/5.0638.3.0.1)",
	// 	BrowserUnknown, 0, PlatformUnknown, OSUnknown, 0, DevicePhone},

	// WebOS reports itself as safari :(
	{"Mozilla/5.0 (webOS/1.0; U; en-US) AppleWebKit/525.27.1 (KHTML, like Gecko) Version/1.0 Safari/525.27.1 Pre/1.0",
		BrowserUnknown, 1, PlatformLinux, OSWebOS, 0, DevicePhone},

	{"Mozilla/5.0 (webOS/1.4.1.1; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		BrowserUnknown, 1, PlatformLinux, OSWebOS, 0, DevicePhone},

	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; de-DE) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
		BrowserUnknown, 0, PlatformLinux, OSWebOS, 0, DeviceTablet},

	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.2; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/234.40.1 Safari/534.6 TouchPad/1.0",
		BrowserUnknown, 0, PlatformLinux, OSWebOS, 0, DeviceTablet},

	{"Opera/9.30 (Nintendo Wii; U; ; 2047-7; fr)",
		BrowserOpera, 9, PlatformNintendo, OSNintendo, 0, DeviceConsole},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US",
		BrowserUnknown, 0, PlatformNintendo, OSNintendo, 0, DeviceConsole},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.6 NintendoBrowser/2.0.0.9362.US",
		BrowserUnknown, 0, PlatformNintendo, OSNintendo, 0, DeviceConsole},

	// TODO fails to get opera first -- but is this a real UA string or an uncommon spoof?
	// {"Mozilla/4.0 (compatible; MSIE 5.0; Windows 2000) Opera 6.0 [en]",
	// 	BrowserIE, 5, PlatformWindows, OSWindows, 4, DeviceComputer},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		BrowserIE, 5, PlatformWindows, OSWindows, 4, DeviceComputer},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		BrowserIE, 7, PlatformWindows, OSWindows, 7, DeviceComputer},

	{"Mozilla/5.0 (Windows; U; Windows NT 6.1; sk; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7",
		BrowserFirefox, 3, PlatformWindows, OSWindows, 7, DeviceComputer},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		BrowserIE, 10, PlatformWindows, OSWindows, 8, DeviceComputer},

	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) YaBrowser/1.0.1084.5402 Chrome/19.0.1084.5402 Safari/536.5",
		BrowserChrome, 19, PlatformWindows, OSWindows, 8, DeviceComputer},

	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.15 (KHTML, like Gecko) Chrome/24.0.1295.0 Safari/537.15",
		BrowserChrome, 24, PlatformWindows, OSWindows, 8, DeviceComputer},

	{"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
		BrowserIE, 11, PlatformWindows, OSWindows, 8, DeviceTablet},

	{"Mozilla/5.0 (IE 11.0; Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		BrowserIE, 11, PlatformWindows, OSWindows, 8, DeviceComputer},

	// {"Mozilla/4.0 (compatible; MSIE 4.01; Windows 95)",
	// 	BrowserIE, 5, PlatformWindows, OSWindows95, 5, DeviceComputer},

	// {"Mozilla/4.0 (compatible; MSIE 5.0; Windows 95) Opera 6.02 [en]",
	// 	BrowserIE, 5, PlatformWindows, OSWindows95, 5, DeviceComputer},

	// {"Mozilla/4.0 (compatible; MSIE 6.0b; Windows 98; YComp 5.0.0.0)",
	// 	BrowserIE, 6, PlatformWindows, OSWindows98, 5, DeviceComputer},

	// {"Mozilla/4.0 (compatible; MSIE 4.01; Windows 98)",
	// 	BrowserIE, 4, PlatformWindows, OSWindows98, 5, DeviceComputer},

	// {"Mozilla/5.0 (Windows; U; Windows 98; en-US; rv:1.8.1.8pre) Gecko/20071019 Firefox/2.0.0.8 Navigator/9.0.0.1",
	// 	BrowserFirefox, 2, PlatformWindows, OSWindows98, 5, DeviceComputer},

	//Can't parse due to limitation of user agent library
	// {"Mozilla/5.0 (Windows; U; Windows CE 5.1; rv:1.8.1a3) Gecko/20060610 Minimo/0.016",
	// 	BrowserUnknown, 0, PlatformWindowsPhone, OSWindowsPhone, 0, DevicePhone},

	// {"Mozilla/4.0 (compatible; MSIE 4.01; Windows CE; 176x220)",
	// 	BrowserIE, 4, PlatformWindowsPhone, OSWindowsPhone, 0, DevicePhone},

	// Can't parse browser due to limitation of user agent library
	// {"Mozilla/4.0 (compatible; MSIE 5.0; Windows ME) Opera 6.0 [de]",
	// 	BrowserUnknown, OSWindowsME, DeviceComputer},

	{"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; InfoPath.2; .NET CLR 3.5.21022; .NET CLR 3.5.30729; MS-RTC LM 8; OfficeLiveConnector.1.4; OfficeLivePatch.1.3; .NET CLR 3.0.30729)",
		BrowserIE, 8, PlatformWindows, OSWindows, 6, DeviceComputer},

	{"Mozilla/5.0 (Windows; U; Windows NT 5.1; cs; rv:1.9.1.8) Gecko/20100202 Firefox/3.5.8",
		BrowserFirefox, 3, PlatformWindows, OSWindows, 5, DeviceComputer},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; )",
		BrowserIE, 7, PlatformWindows, OSWindows, 5, DeviceComputer},

	// Can't parse due to limitation of user agent library
	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Windows Phone 6.5.3.5)",
		BrowserIE, 6, PlatformWindowsPhone, OSWindowsPhone, 6, DevicePhone},

	// desktop mode for Windows Phone 7
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; XBLWP7; ZuneWP7)",
		BrowserIE, 7, PlatformWindows, OSWindows, 7, DeviceComputer},

	// mobile mode for Windows Phone 7
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; HTC; T8788)",
		BrowserIE, 7, PlatformWindowsPhone, OSWindowsPhone, 7, DevicePhone},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0)",
		BrowserIE, 9, PlatformWindowsPhone, OSWindowsPhone, 7, DevicePhone},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 920)",
		BrowserIE, 10, PlatformWindowsPhone, OSWindowsPhone, 8, DevicePhone},

	{"Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch IEMobile/11.0; HTC; Windows Phone 8S by HTC) like Gecko",
		BrowserIE, 11, PlatformWindowsPhone, OSWindowsPhone, 8, DevicePhone},

	{"Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch IEMobile/11.0; NOKIA; 909) like Gecko",
		BrowserIE, 11, PlatformWindowsPhone, OSWindowsPhone, 8, DevicePhone},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)",
		BrowserIE, 9, PlatformXbox, OSXbox, 6, DeviceConsole},
}

func TestAgentSurfer(t *testing.T) {
	//bp := new(BrowserProfile)
	for i, determined := range testUAVars {
		//bp.Parse(determined.UA)
		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse(determined.UA)

		if browserName != determined.browserName {
			t.Errorf("%d browserName: got %v, wanted %v", i, browserName, determined.browserName)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if browserVersion != determined.browserVersion {
			t.Errorf("%d browser version: got %d, wanted %d", i, browserVersion, determined.browserVersion)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if platform != determined.Platform {
			t.Errorf("%d platform: got %v, wanted %v", i, platform, determined.Platform)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if osName != determined.osName {
			t.Errorf("%d os: got %s, wanted %s", i, osName, determined.osName)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if osVersion != determined.osVersion {
			t.Errorf("%d os version: got %d, wanted %d", i, osVersion, determined.osVersion)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if deviceType != determined.deviceType {
			t.Errorf("%d device type: got %v, wanted %v", i, deviceType, determined.deviceType)
			t.Logf("%d agent: %s", i, determined.UA)
		}
	}
}

// func TestExternalFile(t *testing.T) {
// 	// open list of UA strings
// 	inputFile, err := os.Open("./ua_test_set_bs.txt")

// 	if err != nil {
// 		panic(err)
// 		os.Exit(1)
// 	}
// 	defer inputFile.Close()

// 	// prepare yourself
// 	reader := bufio.NewReader(inputFile)
// 	scanner := bufio.NewScanner(reader)

// 	// goodbye console
// 	i := 1
// 	for scanner.Scan() {
// 		fmt.Println("-  ", i, "  -")
// 		fmt.Println(scanner.Text())
// 		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse(scanner.Text())
// 		fmt.Println(browserName, browserVersion, platform, osName, osVersion, deviceType)
// 		i++
// 	}
// }

func BenchmarkAgentSurfer(b *testing.B) {
	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalSystem(b *testing.B) {
	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalSystem(testUAVars[i%num].UA)
	}
}

// // Chrome for Mac
// func BenchmarkParseChromeMac(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36")
// 	}
// }

// // Chrome for Windows
// func BenchmarkParseChromeWin(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.134 Safari/537.36")
// 	}
// }

// // Chrome for Android
// func BenchmarkParseChromeAndroid(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse("Mozilla/5.0 (Linux; Android 4.4.2; GT-P5210 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.93 Safari/537.36")
// 	}
// }

// // Safari for Mac
// func BenchmarkParseSafariMac(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12")
// 	}
// }

// // Safari for iPad
// func BenchmarkParseSafariiPad(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		browserName, browserVersion, platform, osName, osVersion, deviceType, _ := Parse("Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B440 Safari/600.1.4")
// 	}
// }

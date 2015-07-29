package user_agent_surfer

import (
	"testing"
)

var testUAStrings = []struct {
	UA             string
	browserName    string
	browserVersion int
	Platform       string
	osName         string
	osVersion      int
	deviceType     DeviceType
}{

	// iPhone
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		"safari", 6, "iphone", "ios", 7, DevicePhone},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		"safari", 8, "iphone", "ios", 8, DevicePhone},

	// iPad
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		"safari", 4, "ipad", "ios", 3, DeviceTablet},

	{"Mozilla/5.0 (iPad; CPU OS 9_0 like Mac OS X) AppleWebKit/601.1.17 (KHTML, like Gecko) Version/8.0 Mobile/13A175 Safari/600.1.4",
		"safari", 8, "ipad", "ios", 9, DeviceTablet},

	// Chrome
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36",
		"chrome", 43, "mac", "os x", 10, DeviceComputer},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/534.48.3",
		"chrome", 19, "iphone", "ios", 5, DevicePhone},

	// Safari
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		"safari", 8, "mac", "os x", 10, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		"safari", 3, "mac", "os x", 5, DeviceComputer},

	// Firefox
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
		"firefox", 1, "iphone", "ios", 8, DevicePhone},

	{"Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0",
		"firefox", 41, "linux", "android", 4, DeviceTablet},

	{"Mozilla/5.0 (Android; Mobile; rv:40.0) Gecko/40.0 Firefox/40.0",
		"firefox", 40, "linux", "android", 0, DevicePhone},

	// Silk
	{"Mozilla/5.0 (Linux; U; Android 4.4.3; de-de; KFTHWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.47 like Chrome/37.0.2026.117 Safari/537.36",
		"silk", 3, "linux", "android", 4, DeviceTablet},

	{"Mozilla/5.0 (Linux; U; en-us; KFJWI Build/IMM76D) AppleWebKit/535.19 (KHTML like Gecko) Silk/2.4 Safari/535.19 Silk-Acceleratedtrue",
		"silk", 2, "linux", "unknown", 0, DeviceTablet},

	// Opera
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 OPR/18.0.1284.68",
		"opera", 18, "windows", "7", 6, DeviceComputer},

	// Internet Explorer -- https://msdn.microsoft.com/en-us/library/hh869301(v=vs.85).aspx
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		"ie", 10, "windows", "8", 6, DeviceComputer},

	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		"ie", 11, "windows", "8", 6, DeviceComputer},

	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.123",
		"ie", 12, "windows", "10", 10, DeviceComputer},

	{"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; DEVICE INFO) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.123",
		"ie", 12, "windows phone", "10", 10, DevicePhone},

	{"Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 520) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537",
		"ie", 11, "windows phone", "8", 8, DevicePhone},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		"ie", 5, "windows", "2000", 5, DeviceComputer},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		"ie", 7, "windows", "7", 6, DeviceComputer},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)", //Windows Surface RT tablet
		"ie", 10, "windows", "8", 6, DeviceTablet},

	// UC Browser
	{"Mozilla/5.0 (Linux; U; Android 2.3.4; en-US; MT11i Build/4.0.2.A.0.62) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.0.1.275 U3/0.8.0 Mobile Safari/534.31",
		"ucbrowser", 9, "linux", "android", 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.0.4; en-US; Micromax P255 Build/IMM76D) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.2.0.308 U3/0.8.0 Mobile Safari/534.31",
		"ucbrowser", 9, "linux", "android", 4, DevicePhone},

	{"UCWEB/2.0 (Java; U; MIDP-2.0; en-US; MicromaxQ5) U2/1.0.0 UCBrowser/9.4.0.342 U2/1.0.0 Mobile",
		"ucbrowser", 9, "unknown", "unknown", 0, DevicePhone},

	// Nokia Browser
	// {"Mozilla/5.0 (Series40; Nokia501/14.0.4/java_runtime_version=Nokia_Asha_1_2; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/4.0.0.0.45",
	// 	"nokia", 4, "unknown", "unknown", 0, DevicePhone},

	// {"Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaN8-00/111.040.1511; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/535.1 (KHTML, like Gecko) NokiaBrowser/8.3.1.4 Mobile Safari/535.1",
	// 	"nokia", 8, "unknown", "unknown", 0, DevicePhone},

	// {"NokiaN97/21.1.107 (SymbianOS/9.4; Series60/5.0 Mozilla/5.0; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebkit/525 (KHTML, like Gecko) BrowserNG/7.1.4",
	// 	"nokia", 7, "unknown", "unknown", 0, DevicePhone},

	// ChromeOS
	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		"chrome", 4, "linux", "chromeos", 0, DeviceComputer},

	// WebOS
	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; de-DE) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
		"unknown", 0, "linux", "webos", 0, DeviceTablet},

	{"Mozilla/5.0 (webOS/1.4.1.1; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		"unknown", 1, "linux", "webos", 0, DeviceUnknown},

	// Android WebView (Android <= 4.3)
	{"Mozilla/5.0 (Linux; U; Android 2.2; en-us; DROID2 GLOBAL Build/S273) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"android", 4, "linux", "android", 2, DevicePhone},

	{"Mozilla/5.0 (Linux; U; Android 4.0.3; de-ch; HTC Sensation Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"android", 4, "linux", "android", 4, DevicePhone},

	// BlackBerry
	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+",
		"blackberry", 7, "blackberry", "blackberry", 0, DeviceTablet},

	{"Mozilla/5.0 (BB10; Kbd) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.2.1.1925 Mobile Safari/537.35+",
		"blackberry", 10, "blackberry", "blackberry", 0, DevicePhone},

	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) BlackBerry8703e/4.1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/104",
		"blackberry", 0, "blackberry", "blackberry", 0, DevicePhone},

	// Windows Phone
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 625; ANZ941)",
		"ie", 10, "windows phone", "8", 8, DevicePhone},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 900)",
		"ie", 9, "windows phone", "7", 7, DevicePhone},

	// Kindle
	{"Mozilla/5.0 (Linux; U; en-US) AppleWebKit/528.5+ (KHTML, like Gecko, Safari/528.5+) Version/4.0 Kindle/3.0 (screen 600×800; rotate)",
		"unknown", 4, "kindle", "kindle", 0, DeviceTablet},

	{"Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+",
		"unknown", 5, "kindle", "kindle", 0, DeviceTablet},

	// Nintendo
	{"Opera/9.30 (Nintendo Wii; U; ; 2047-7; fr)",
		"opera", 9, "nintendo", "nintendo", 0, DeviceConsole},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US",
		"unknown", 0, "nintendo", "nintendo", 0, DeviceConsole},

	// Xbox
	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)", //Xbox 360
		"ie", 9, "xbox", "xbox", 6, DeviceConsole},

	// Playstation

	{"Mozilla/5.0 (Playstation Vita 1.61) AppleWebKit/531.22.8 (KHTML, like Gecko) Silk/3.2",
		"silk", 3, "playstation", "playstation", 0, DeviceConsole},

	// Smart TVs and TV dongles
	{"Mozilla/5.0 (CrKey armv7l 1.4.15250) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.0 Safari/537.36", // Chromecast
		"chrome", 31, "unknown", "unknown", 0, DeviceTV},

	{"Mozilla/5.0 (Linux; GoogleTV 3.2; VAP430 Build/MASTER) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.77 Safari/534.24", // Google TV
		"chrome", 11, "linux", "unknown", 0, DeviceTV},

	{"Mozilla/5.0 (Linux; Android 5.0; ADT-1 Build/LPX13D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Mobile Safari/537.36", // Android TV
		"chrome", 40, "linux", "android", 5, DeviceTV},

	{"Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22", // Amazon Fire
		"chrome", 25, "linux", "android", 4, DeviceTV},

	{"Mozilla/5.0 (Unknown; Linux armv7l) AppleWebKit/537.1+ (KHTML, like Gecko) Safari/537.1+ LG Browser/6.00.00(+mouse+3D+SCREEN+TUNER; LGE; GLOBAL-PLAT5; 03.07.01; 0x00000001;); LG NetCast.TV-2013/03.17.01 (LG, GLOBAL-PLAT4, wired)", // LG TV
		"unknown", 0, "linux", "unknown", 0, DeviceTV},

	{"Mozilla/5.0 (X11; FreeBSD; U; Viera; de-DE) AppleWebKit/537.11 (KHTML, like Gecko) Viera/3.10.0 Chrome/23.0.1271.97 Safari/537.11", // Panasonic Viera
		"chrome", 23, "linux", "linux", 0, DeviceTV},

	{"Mozilla/5.0 (DTV) AppleWebKit/531.2+ (KHTML, like Gecko) Espial/6.1.5 AQUOSBrowser/2.0 (US01DTV;V;0001;0001)", // Sharp Aquos
		"unknown", 0, "unknown", "unknown", 0, DeviceTV},

	{"Roku/DVP-5.2 (025.02E03197A)", // Roku
		"unknown", 0, "unknown", "unknown", 0, DeviceTV},

	// Bots
	// {"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	// 	"bot", 0, "bot", "bot", 0, DeviceBot},

	// {"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	// 	"bot", 0, "bot", "bot", 0, DeviceBot},

	// Unknown or partially handled
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3", //Seamonkey (~FF)
		"firefox", 0, "mac", "os x", 4, DeviceComputer},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre", //Camino (~FF)
		"unknown", 0, "mac", "os x", 5, DeviceComputer},

	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", //firefox OS
		"firefox", 26, "unknown", "unknown", 0, DevicePhone},

	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.45 Safari/535.19", //chrome for android having requested desktop site
		"chrome", 18, "linux", "linux", 0, DeviceComputer},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		"opera", 10, "unknown", "unknown", 0, DevicePhone},
}

func TestAgentSurfer(t *testing.T) {
	bp := new(BrowserProfile)
	for i, determined := range testUAStrings {
		bp.Parse(determined.UA)

		if bp.Browser.Name != determined.browserName {
			t.Errorf("%d browserName: got %s, wanted %s", i, bp.Browser.Name, determined.browserName)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if bp.Browser.Version != determined.browserVersion {
			t.Errorf("%d browser version: got %d, wanted %d", i, bp.Browser.Version, determined.browserVersion)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if bp.Platform != determined.Platform {
			t.Errorf("%d platform: got %s, wanted %s", i, bp.Platform, determined.Platform)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if bp.OS.Name != determined.osName {
			t.Errorf("%d os: got %s, wanted %s", i, bp.OS.Name, determined.osName)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if bp.OS.Version != determined.osVersion {
			t.Errorf("%d os version: got %d, wanted %d", i, bp.OS.Version, determined.osVersion)
			t.Logf("%d agent: %s", i, determined.UA)
		}

		if bp.DeviceType != determined.deviceType {
			t.Errorf("%d device type: got %v, wanted %v", i, bp.DeviceType, determined.deviceType)
			t.Logf("%d agent: %s", i, determined.UA)
		}
	}
}

func BenchmarkAgentSurfer(b *testing.B) {
	num := len(testUAStrings)
	bp := new(BrowserProfile)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bp.Parse(testUAStrings[i%num].UA)
	}
}

func BenchmarkFullParse(b *testing.B) {
	bp := new(BrowserProfile)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bp.Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36")
	}
}

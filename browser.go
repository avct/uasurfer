package uasurfer

import (
	"regexp"
	//"strconv"
	"strings"
)

var (
	safariFingerprints = regexp.MustCompile("\\w{3}\\/\\d")
	bVersion           = regexp.MustCompile("version/\\d+")
	chromeVersion      = regexp.MustCompile("(chrome|crios)/\\d+")
	ieVersion          = regexp.MustCompile("(msie\\s|edge/)\\d+")
	tridentVersion     = regexp.MustCompile("trident/\\d+")
	firefoxVersion     = regexp.MustCompile("(firefox|fxios)/\\d+")
	ucVersion          = regexp.MustCompile("ucbrowser/\\d+")
	oprVersion         = regexp.MustCompile("(opr|opios)/\\d+")
	operaVersion       = regexp.MustCompile("opera/\\d+")
	silkVersion        = regexp.MustCompile("silk/\\d+")
	spotifyVersion     = regexp.MustCompile("spotify/\\d+")
)

// Browser struct contains the lowercase name of the browser, along
// with its major browser version number. Browser are grouped together without
// consideration for device. For example, Chrome (Chrome/43.0) and Chrome for iOS
// (CriOS/43.0) would both return as "chrome" (name) and 43 (version). Similarly
// Internet Explorer 11 and Edge 12 would return as "ie" and "11" or "12", respectively.
// type Browser struct {
// 	Name    BrowserName
// 	Version int
// }

// Retrieve the espoused major version of the browser if possible, prioritizing
// the "Version/#" UA attribute over others. Set to 0 if no version
// is obtainable. A lowercase browser name (string) and its
// version (int) is returned.
func evalBrowserName(ua string) BrowserName {

	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") || strings.Contains(ua, "bb10") { //blackberry goes first because it reads as MSIE & Safari
		return BrowserBlackberry
	}

	if strings.Contains(ua, "applewebkit") {

		if strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/") {
			return BrowserOpera
		}

		if strings.Contains(ua, "silk/") {
			return BrowserSilk
		}

		if strings.Contains(ua, "edge/") || strings.Contains(ua, "iemobile/") || strings.Contains(ua, "msie ") {
			return BrowserIE
		}

		if strings.Contains(ua, "ucbrowser/") || strings.Contains(ua, "ucweb/") {
			return BrowserUCBrowser
		}

		//Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
		if strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "chromium/") {
			return BrowserChrome
		}

		if strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android") {
			// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
			return BrowserAndroid
		}

		if strings.Contains(ua, "fxios") {
			return BrowserFirefox
		}

		if strings.Contains(ua, " spotify/") {
			return BrowserSpotify
		}

		if strings.Contains(ua, "like gecko") {
			// Safari is the most generic, archtypical User-Agent on the market -- it's identified by making sure effectively by checking for attribute purity. It's fingerprint should have 4 or 5 total x/y attributes, 'mobile/version' being optional
			safariFingerprints := len(safariFingerprints.FindAllString(ua, -1))

			if (safariFingerprints == 4 || safariFingerprints == 5) && strings.Contains(ua, "version/") && strings.Contains(ua, "safari/") && strings.Contains(ua, "mozilla/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") {
				return BrowserSafari
			}
		}

		// Google's search app on iPhone, leverages native Safari rather than Chrome
		if strings.Contains(ua, " gsa/") {
			return BrowserSafari
		}
	}

	if strings.Contains(ua, "msie") || strings.Contains(ua, "trident") {
		return BrowserIE
	}

	if strings.Contains(ua, "gecko") {
		if strings.Contains(ua, "firefox") || strings.Contains(ua, "iceweasel") || strings.Contains(ua, "seamonkey") || strings.Contains(ua, "icecat") {
			return BrowserFirefox
		}
	}

	if strings.Contains(ua, "presto") || strings.Contains(ua, "opera") {
		return BrowserOpera
	}

	if strings.Contains(ua, "ucbrowser") {
		return BrowserUCBrowser
	}

	return BrowserUnknown
}

func evalBrowserVersion(ua string, browserName BrowserName) int {
	// Find browser version using 3 methods in order:
	// 1st: look for generic version/#
	// 2nd: look for browser-specific instructions (e.g. chrome/34)
	// 3rd: infer from OS (iOS only)

	// if there is a 'version/#' attribute with numeric version, use it -- except for Chrome since Android vendors sometimes hijack version/#
	if browserName != BrowserChrome && bVersion.MatchString(ua) {
		v := bVersion.FindString(ua)
		v = strings.Split(v, "/")[1]
		i := strToInt(v)
		return i
	}

	switch browserName {
	case BrowserChrome:
		// match both chrome and crios
		return getMajorVersion(ua, chromeVersion)
	case BrowserIE:
		if strings.Contains(ua, "msie") || strings.Contains(ua, "edge") {
			i := getMajorVersion(ua, ieVersion)

			return i
		}

		// switch based on trident version indicator https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if strings.Contains(ua, "trident") {
			i := getMajorVersion(ua, tridentVersion)
			switch i {
			case 3:
				return 7
			case 4:
				return 8
			case 5:
				return 9
			case 6:
				return 10
			case 7:
				return 11
			}
		}

		return 0

	case BrowserFirefox:
		i := getMajorVersion(ua, firefoxVersion)

		return i

	case BrowserSafari: // executes typically if we're on iOS and not using a familiar browser
		i := getiOSSafariVersion(ua)

		return i

	case BrowserUCBrowser:
		i := getMajorVersion(ua, ucVersion)

		return i
	case BrowserOpera:
		if strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/") {
			i := getMajorVersion(ua, oprVersion)
			return i
		}

		i := getMajorVersion(ua, operaVersion)
		return i
	case BrowserSilk:
		i := getMajorVersion(ua, silkVersion)
		return i

	case BrowserSpotify:
		i := getMajorVersion(ua, spotifyVersion)
		return i

	default:
		return 0
	}

	// Handle no match
	return 0
}

// Subfunction of evalBrowser() that takes two parameters: regex (string) and
// user agent (string) and returns the number as a string. '0' denotes no version.
func getMajorVersion(ua string, browserVersion *regexp.Regexp) int {

	ver := browserVersion.FindString(ua)

	if ver != "" {
		if strings.Contains(ver, "/") {
			ver = strings.Split(ver, "/")[1] //e.g. "version/10.0.2"
			i := strToInt(ver)
			return i
		}

		if strings.Contains(ver, " ") {
			ver = strings.Split(ver, " ")[1] //e.g. "msie 10.0"
			i := strToInt(ver)
			return i
		}
		return 0
	}

	i := strToInt(ver)

	return i
}

// getiOSSafariVersion accepts a full UA string and returns
// an `int` of the major version of Safari. The latest browser
// version released for the OS is used. This function is used
// in uncommon scenarios such as the Google search app browser
func getiOSSafariVersion(ua string) int {
	v := getiOSVersion(ua)

	switch v {
	case 1: // OS version
		return 2 // Safari version
	case 2:
		return 3
	case 3:
		return 4
	case 4:
		return 4
	case 5:
		return 5
	case 6:
		return 6
	case 7:
		return 6
	case 8:
		return 8
	case 9:
		return 8
	}

	return 0 // default
}

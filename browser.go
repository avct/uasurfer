package uasurfer

import (
	"regexp"
	//"strconv"
	"strings"
)

var (
	// safariFingerprints = regexp.MustCompile("\\w{3}\\/\\d")
	bVersion       = regexp.MustCompile("version/\\d+") // standard browser versioning e.g. "Version/10.0"
	chromeVersion  = regexp.MustCompile("(chrome|crios|crmo)/\\d+")
	ieVersion      = regexp.MustCompile("(msie\\s|edge/)\\d+")
	tridentVersion = regexp.MustCompile("trident/\\d+")
	firefoxVersion = regexp.MustCompile("(firefox|fxios)/\\d+")
	ucVersion      = regexp.MustCompile("ucbrowser/\\d+")
	oprVersion     = regexp.MustCompile("(opr|opios)/\\d+")
	operaVersion   = regexp.MustCompile("opera/\\d+")
	silkVersion    = regexp.MustCompile("silk/\\d+")
	spotifyVersion = regexp.MustCompile("spotify/\\d+")
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

	// Blackberry goes first because it reads as MSIE & Safari
	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") || strings.Contains(ua, "bb10") || strings.Contains(ua, "rim ") {
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

		// Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
		if strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "chromium/") || strings.Contains(ua, "crmo/") {
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
			// presume it's safari unless an esoteric browser is being specified (webOSBrowser, SamsungBrowser, etc.)
			if strings.Contains(ua, "mozilla/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") && strings.Contains(ua, "safari/") && !strings.Contains(ua, "browser/") && !strings.Contains(ua, "os/") {

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

	if strings.Contains(ua, "phantomjs") || strings.Contains(ua, "googlebot") {
		return BrowserBot
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
		ua = bVersion.FindString(ua)
		s := strings.Index(ua, "/")
		ua = ua[s+1:]
		i := strToInt(ua)
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

		// get MSIE version from trident version https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if strings.Contains(ua, "trident") {
			i := getMajorVersion(ua, tridentVersion)
			// convert trident versions 3-7 to MSIE version
			if (i >= 3) && (i <= 7) {
				return i + 4
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
}

// Subfunction of evalBrowser() that takes two parameters: regex (string) and
// user agent (string) and returns the number as a string. '0' denotes no version.
func getMajorVersion(ua string, browserVersion *regexp.Regexp) int {

	ver := browserVersion.FindString(ua)

	if ver != "" {
		if strings.Contains(ver, "/") {
			return getVersionNumber(ver, "/")
		}
		return getVersionNumber(ver, " ")
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

	// early Safari used a version number +1 to OS version
	if (v <= 3) && (v >= 1) {
		return v + 1
	}

	// later Safari typically matches iOS version after v3
	if v >= 4 {
		return v
	}

	return 0 // default
}

// returns version number to the left of a delimiter
func getVersionNumber(s, delim string) int {
	ind := strings.Index(s, delim)
	i := strToInt(s[ind+1:])
	return i
}

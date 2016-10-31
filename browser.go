package uasurfer

import (
	"regexp"
	//"strconv"
	"strings"
)

var (
	bVersion       = regexp.MustCompile("version/\\d+(?:[_\\.]\\d+)*") // standard browser versioning e.g. "Version/10.0"
	chromeVersion  = regexp.MustCompile("(chrome|crios|crmo)/\\d+(?:\\.\\d+)*")
	ieVersion      = regexp.MustCompile("(msie\\s|edge/)\\d+(?:\\.\\d+)*")
	tridentVersion = regexp.MustCompile("trident/\\d+(?:\\.\\d+)*")
	firefoxVersion = regexp.MustCompile("(firefox|fxios)/\\d+(?:\\.\\d+)*")
	ucVersion      = regexp.MustCompile("ucbrowser/\\d+(?:\\.\\d+)*")
	oprVersion     = regexp.MustCompile("(opr|opios)/\\d+(?:\\.\\d+)*")
	operaVersion   = regexp.MustCompile("opera/\\d+(?:\\.\\d+)*")
	silkVersion    = regexp.MustCompile("silk/\\d+(?:\\.\\d+)*")
	spotifyVersion = regexp.MustCompile("spotify/\\d+(?:\\.\\d+)*")
)

// Browser struct contains the lowercase name of the browser, along
// with its browser version number. Browser are grouped together without
// consideration for device. For example, Chrome (Chrome/43.0) and Chrome for iOS
// (CriOS/43.0) would both return as "chrome" (name) and 43.0 (version). Similarly
// Internet Explorer 11 and Edge 12 would return as "ie" and "11" or "12", respectively.
// type Browser struct {
// 		Name    BrowserName
// 		Version struct {
// 			Major int
// 			Minor int
// 			Patch int
// 		}
// }

// Retrieve browser name from UA strings
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

// Retrieve browser version returning 0.0.0 if none is available
// Methods used in order:
// 1st: look for generic version/#
// 2nd: look for browser-specific instructions (e.g. chrome/34)
// 3rd: infer from OS (iOS only)
func evalBrowserVersion(ua string, browserName BrowserName) Version {

	// if there is a 'version/#' attribute with numeric version, use it -- except for Chrome since Android vendors sometimes hijack version/#
	if browserName != BrowserChrome && bVersion.MatchString(ua) {
		ua = bVersion.FindString(ua)
		s := strings.Index(ua, "/")
		ua = ua[s+1:]
		i := strToVersion(ua)
		return i
	}

	switch browserName {
	case BrowserChrome:
		// match both chrome and crios
		return getVersion(ua, chromeVersion)
	case BrowserIE:
		if strings.Contains(ua, "msie") || strings.Contains(ua, "edge") {
			return getVersion(ua, ieVersion)
		}

		// get MSIE version from trident version https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if strings.Contains(ua, "trident") {
			i := getVersion(ua, tridentVersion)
			// convert trident versions 3-7 to MSIE version
			if (i.Major >= 3) && (i.Major <= 7) {
				i.Major += 4
				return i
			}
		}

		return Version{}

	case BrowserFirefox:
		return getVersion(ua, firefoxVersion)

	case BrowserSafari: // executes typically if we're on iOS and not using a familiar browser
		return getiOSSafariVersion(ua)

	case BrowserUCBrowser:
		return getVersion(ua, ucVersion)

	case BrowserOpera:
		if strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/") {
			return getVersion(ua, oprVersion)
		}

		return getVersion(ua, operaVersion)

	case BrowserSilk:
		return getVersion(ua, silkVersion)

	case BrowserSpotify:
		return getVersion(ua, spotifyVersion)

	default:
		return Version{}
	}
}

// Subfunction of evalBrowser() that takes two parameters: regex (string) and
// user agent (string) and returns Version. '0.0.0' denotes no version.
func getVersion(ua string, browserVersion *regexp.Regexp) Version {

	ver := browserVersion.FindString(ua)

	if ver != "" {
		if strings.Contains(ver, "/") {
			return getVersionNumber(ver, "/")
		}
		return getVersionNumber(ver, " ")
	}

	return strToVersion(ver)
}

// getiOSSafariVersion accepts a full UA string and returns
// an Version of Safari. The latest browser
// version released for the OS is used. This function is used
// in uncommon scenarios such as the Google search app browser
func getiOSSafariVersion(ua string) Version {
	v := getiOSVersion(ua)

	// early Safari used a version number +1 to OS version
	if (v.Major <= 3) && (v.Major >= 1) {
		v.Major++
	}

	return v
}

// returns version number to the left of a delimiter
func getVersionNumber(s, delim string) Version {
	ind := strings.Index(s, delim)
	return strToVersion(s[ind+1:])
}

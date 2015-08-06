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
	oprVersion         = regexp.MustCompile("opr/\\d+")
	operaVersion       = regexp.MustCompile("opera/\\d+")
	silkVersion        = regexp.MustCompile("silk/\\d+")
)

// Browser struct contains the lowercase name of the browser, along
// with its major browser version number. Browser are grouped together without
// consideration for device. For example, Chrome (Chrome/43.0) and Chrome for iOS
// (CriOS/43.0) would both return as "chrome" (name) and 43 (version). Similarly
// Internet Explorer 11 and Edge 12 would return as "ie" and "11" or "12", respectively.
type Browser struct {
	Name    BrowserName
	Version int
}

// Retrieve the espoused major version of the browser if possible, prioritizing
// the "Version/#" UA attribute over others. Set to 0 if no version
// is obtainable. A lowercase browser name (string) and its
// version (int) is returned.
func (b *BrowserProfile) evalBrowserName(ua string) BrowserName {

	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") || strings.Contains(ua, "bb10") { //blackberry goes first because it reads as MSIE & Safari
		return BrowserBlackberry
	}

	if strings.Contains(ua, "applewebkit") {

		if strings.Contains(ua, "opr/") {
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

		if strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "chromium/") { //Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
			return BrowserChrome
		}

		if strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android") {
			// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
			return BrowserAndroid
		}

		if strings.Contains(ua, "fxios") {
			return BrowserFirefox
		}

		if strings.Contains(ua, "like gecko") {
			// Safari is the most generic, archtypical User-Agent on the market -- it's identified by making sure effectively by checking for attribute purity. It's fingerprint should have 4 or 5 total x/y attributes, 'mobile/version' being optional
			safariFingerprints := len(safariFingerprints.FindAllString(ua, -1))

			if (safariFingerprints == 4 || safariFingerprints == 5) && strings.Contains(ua, "version/") && strings.Contains(ua, "safari/") && strings.Contains(ua, "mozilla/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") {
				return BrowserSafari
			}
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

	// if strings.Contains(ua, "nintendo") {
	// 	return BrowserNintendo
	// }

	return BrowserUnknown
}

func (b *BrowserProfile) evalBrowserVersion(ua string) string {
	// Find browser version using 3 methods in order:
	// 1st: look for generic version/#
	// 2nd: look for browser-specific instructions (e.g. chrome/34)
	// 3rd: infer from OS
	v := ""

	// if there is a 'version/#' attribute with numeric version, use it
	if bVersion.MatchString(ua) {
		v = bVersion.FindString(ua)
		return strings.Split(v, "/")[1]
	}

	switch b.Browser.Name {
	case BrowserChrome:
		// match both chrome and crios
		return getMajorVersion(ua, chromeVersion)
	case BrowserIE:
		if strings.Contains(ua, "msie") || strings.Contains(ua, "edge") {
			return getMajorVersion(ua, ieVersion)
		}

		// switch based on trident version indicator https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if strings.Contains(ua, "trident") {
			v = getMajorVersion(ua, tridentVersion)
			switch v {
			case "3":
				return "7"
			case "4":
				return "8"
			case "5":
				return "9"
			case "6":
				return "10"
			case "7":
				return "11"
			}
		}

		return "0"

	case BrowserFirefox:
		return getMajorVersion(ua, firefoxVersion)
	case BrowserUCBrowser:
		return getMajorVersion(ua, ucVersion)
	case BrowserOpera:
		if strings.Contains(ua, "opr/") {
			return getMajorVersion(ua, oprVersion)
		}

		return getMajorVersion(ua, operaVersion)
	case BrowserSilk:
		return getMajorVersion(ua, silkVersion)
	default:
		return "0"
	}

	// backup plans if we still don't know the version: guestimate based on highest available to the device
	// if v == "0" {
	// 	switch b.OS.Name {
	// 		case "iOS":
	// 			switch os.version {
	// 				case 1:
	// 				case 2:
	// 				case 3:
	// 				case 4:
	// 				case 5:
	// 				case 6:
	// 				case 7:
	// 				case 8:
	// 				case 9:
	// 				case 10:
	// 		case "Android":
	// 			switch os.version {
	// 				case 1:
	// 				case 2:
	// 				case 3:
	// 				case 4:
	// 				case 5:
	// 				case 6:
	// 				case 7:
	// 				case 8:
	// 				case 9:
	// 				case 10:
	// 		case "OS X":
	// 			switch os.version {
	// 				case 1:
	// 				case 2:
	// 				case 3:
	// 				case 4:
	// 				case 5:
	// 				case 6:
	// 				case 7:
	// 				case 8:
	// 				case 9:
	// 				case 10:
	// 				case 11:
	// 	}
	// }

	// Handle no match
	return "0"
}

// Subfunction of evalBrowser() that takes two parameters: regex (string) and
// user agent (string) and returns the number as a string. "0" denotes no version.
func getMajorVersion(ua string, bVersion *regexp.Regexp) string {
	ver := bVersion.FindString(ua)

	if ver != "" {
		if strings.Contains(ver, "/") {
			ver = strings.Split(ver, "/")[1] //e.g. "version/10.0.2"
		} else if strings.Contains(ver, " ") {
			ver = strings.Split(ver, " ")[1] //e.g. "msie 10.0"
		} else {
			ver = "0"
		}
	}

	return ver
}

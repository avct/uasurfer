package user_agent_surfer

import (
	"regexp"
	"strconv"
	"strings"
)

// Browser struct contains the lowercase name of the browser, along
// with its major browser version number. Browser are grouped together without
// consideration for device. For example, Chrome (Chrome/43.0) and Chrome for iOS
// (CriOS/43.0) would both return as "chrome" (name) and 43 (version). Similarly
// Internet Explorer 11 and Edge 12 would return as "ie" and "11" or "12", respectively.
type Browser struct {
	Name    string
	Version int
}

// Retrieve the espoused major version of the browser if possible, prioritizing
// the "Version/#" UA attribute over others. Set to 0 if no version
// is obtainable. A lowercase browser name (string) and its
// version (int) is returned.
func (b *BrowserProfile) evalBrowser(ua string) (string, int) {

	// Narrow browser by engine, then inference from other key words
	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") || strings.Contains(ua, "bb10") { //blackberry goes first because it reads as MSIE & Safari really well
		b.Browser.Name = "blackberry"
	} else if strings.Contains(ua, "applewebkit") {
		if strings.Contains(ua, "opr/") {
			b.Browser.Name = "opera"
		} else if strings.Contains(ua, "silk/") {
			b.Browser.Name = "silk"
		} else if strings.Contains(ua, "edge/") || strings.Contains(ua, "iemobile/") || strings.Contains(ua, "msie ") {
			b.Browser.Name = "ie"
		} else if strings.Contains(ua, "ucbrowser/") || strings.Contains(ua, "ucweb/") {
			b.Browser.Name = "ucbrowser"
		} else if strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "chromium/") { //Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
			b.Browser.Name = "chrome"
		} else if strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") {
			// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
			b.Browser.Name = "android"
		} else if strings.Contains(ua, "fxios") {
			b.Browser.Name = "firefox"
		} else if strings.Contains(ua, "like gecko") {
			// Safari is the most generic, archtypical User-Agent on the market -- it's identified by making sure effectively by checking for attribute purity. It's fingerprint should have 4 or 5 total x/y attributes, 'mobile/version' being optional
			safariId, _ := regexp.Compile("\\w{3}\\/\\d")
			safariFingerprint := len(safariId.FindAllString(ua, -1))

			if (safariFingerprint == 4 || safariFingerprint == 5) && strings.Contains(ua, "version/") && strings.Contains(ua, "safari/") && strings.Contains(ua, "mozilla/") {
				b.Browser.Name = "safari"
			}
		}
	} else if strings.Contains(ua, "msie") || strings.Contains(ua, "trident") {
		b.Browser.Name = "ie"
	} else if strings.Contains(ua, "gecko") {
		if strings.Contains(ua, "firefox") || strings.Contains(ua, "iceweasel") || strings.Contains(ua, "seamonkey") || strings.Contains(ua, "icecat") {
			b.Browser.Name = "firefox"
		}
	} else if strings.Contains(ua, "presto") || strings.Contains(ua, "opera") {
		b.Browser.Name = "opera"
	} else if strings.Contains(ua, "nintendo") {
		b.Browser.Name = "nintendo"
	}
	// Find browser version using 3 methods in order:
	// 1st: look for generic version/#
	// 2nd: look for browser-specific instructions (e.g. chrome/34)
	// 3rd: infer from OS
	v := ""
	bVersion, _ := regexp.Compile("version/\\d+")

	// if there is a 'version/#' attribute with numeric version, use it
	if bVersion.MatchString(ua) {
		v = bVersion.FindString(ua)
		v = strings.Split(v, "/")[1]
	} else {
		switch b.Browser.Name {
		case "chrome":
			// match both chrome and crios
			v = getMajorVersion(ua, "(chrome|crios)/\\d+")
		case "ie":
			if strings.Contains(ua, "msie") || strings.Contains(ua, "edge") {
				v = getMajorVersion(ua, "(msie\\s|edge/)\\d+")
			} else {
				// switch based on trident version indicator https://en.wikipedia.org/wiki/Trident_(layout_engine)
				if strings.Contains(ua, "trident") {
					v = getMajorVersion(ua, "trident/\\d+")
					switch v {
					case "3":
						v = "7"
					case "4":
						v = "8"
					case "5":
						v = "9"
					case "6":
						v = "10"
					case "7":
						v = "11"
					}
				}
			}
			if v == "" {
				v = "0"
			}
		case "firefox":
			v = getMajorVersion(ua, "(firefox|fxios)/\\d+")
		case "ucbrowser":
			v = getMajorVersion(ua, "ucbrowser/\\d+")
		case "opera":
			if strings.Contains(ua, "opr/") {
				v = getMajorVersion(ua, "opr/\\d+")
			} else {
				v = getMajorVersion(ua, "opera/\\d+")
			}
		case "silk":
			v = getMajorVersion(ua, "silk/\\d+")
		case "nintendo":
			v = "0" //getMajorVersion(ua, "nintendobrowser/\\d+")
		//case "opera":
		// could be either version/x or opr/x
		default:
			v = "0"
		}
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
	if v == "" {
		v = "0"
	}
	if b.Browser.Name == "" {
		b.Browser.Name = "unknown"
	}

	b.Browser.Version, _ = strconv.Atoi(v)

	return b.Browser.Name, b.Browser.Version
}

// Subfunction of evalBrowser() that takes two parameters: regex (string) and
// user agent (string) and returns the number as a string. "0" denotes no version.
func getMajorVersion(ua string, match string) string {
	bVersion, _ := regexp.Compile(match)
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

// Package uasurfer provides fast and reliable abstraction
// of HTTP User-Agent strings. The philosophy is to identify
// technologies that holds >1% market share, and to avoid
// expending resources and accuracy on guessing at esoteric UA
// strings.
package uasurfer

import "strings"

//go:generate stringer -type=DeviceID,BrowserID,OSID,Platform -output=const_string.go

// DeviceID (uint8) returns a constant.
type DeviceID int8

// A complete list of supported devices in the
// form of constants.
const (
	DeviceUnknown  DeviceID = 0
	DeviceComputer DeviceID = 1
	DeviceTablet   DeviceID = 2
	DevicePhone    DeviceID = 3
	DeviceConsole  DeviceID = 4
	DeviceWearable DeviceID = 5
	DeviceTV       DeviceID = 6
)

var deviceNames = map[DeviceID]string{
	DeviceUnknown:  "Unknown",
	DeviceComputer: "Computer",
	DeviceTablet:   "Tablet",
	DevicePhone:    "Mobile",
	DeviceConsole:  "Console",
	DeviceWearable: "Wearable",
	DeviceTV:       "TV",
}

func (d DeviceID) String() string {
	return deviceNames[d]
}

// BrowserID (uint8) returns a constant.
type BrowserID uint8

// A complete list of supported web browsers in the
// form of constants.
const (
	BrowserUnknown       BrowserID = 0
	BrowserChrome        BrowserID = 1
	BrowserIE            BrowserID = 3
	BrowserSafari        BrowserID = 4
	BrowserFirefox       BrowserID = 5
	BrowserAndroid       BrowserID = 6
	BrowserOpera         BrowserID = 7
	BrowserBlackberry    BrowserID = 8
	BrowserUCBrowser     BrowserID = 9
	BrowserSilk          BrowserID = 10
	BrowserNokia         BrowserID = 11
	BrowserNetFront      BrowserID = 12
	BrowserQQ            BrowserID = 13
	BrowserMaxthon       BrowserID = 14
	BrowserSogouExplorer BrowserID = 15
	BrowserSpotify       BrowserID = 16
	BrowserNintendo      BrowserID = 17
	BrowserSamsung       BrowserID = 18
	BrowserYandex        BrowserID = 19
	BrowserCocCoc        BrowserID = 20
	BrowserBot           BrowserID = 21 // Bot list begins here
	BrowserAppleBot      BrowserID = 22
	BrowserBaiduBot      BrowserID = 23
	BrowserBingBot       BrowserID = 24
	BrowserDuckDuckGoBot BrowserID = 25
	BrowserFacebookBot   BrowserID = 26
	BrowserGoogleBot     BrowserID = 27
	BrowserGoogleAdsBot  BrowserID = 28
	BrowserLinkedInBot   BrowserID = 29
	BrowserMsnBot        BrowserID = 30
	BrowserPingdomBot    BrowserID = 31
	BrowserTwitterBot    BrowserID = 32
	BrowserYandexBot     BrowserID = 33
	BrowserCocCocBot     BrowserID = 34
	BrowserPinterestBot  BrowserID = 35
	BrowserSlackBot      BrowserID = 36
	BrowserSeekportBot   BrowserID = 37
	BrowserYahooBot      BrowserID = 38 // Bot list ends here
)

var browserNames = map[BrowserID]string{
	BrowserUnknown:       "Unknown",
	BrowserChrome:        "Chrome",
	BrowserIE:            "IE",
	BrowserSafari:        "Safari",
	BrowserFirefox:       "Firefox",
	BrowserAndroid:       "Android",
	BrowserOpera:         "Opera",
	BrowserBlackberry:    "Blackberry",
	BrowserUCBrowser:     "UCBrowser",
	BrowserSilk:          "Silk",
	BrowserNokia:         "Nokia",
	BrowserNetFront:      "NetFront",
	BrowserQQ:            "QQ",
	BrowserMaxthon:       "Maxthon",
	BrowserSogouExplorer: "SogouExplorer",
	BrowserSpotify:       "Spotify",
	BrowserNintendo:      "Nintendo",
	BrowserSamsung:       "Samsung",
	BrowserYandex:        "Yandex",
	BrowserCocCoc:        "CocCoc",
	BrowserBot:           "Bot",
	BrowserAppleBot:      "AppleBot",
	BrowserBaiduBot:      "BaiduBot",
	BrowserBingBot:       "BingBot",
	BrowserDuckDuckGoBot: "DuckDuckGoBot",
	BrowserFacebookBot:   "FacebookBot",
	BrowserGoogleBot:     "GoogleBot",
	BrowserGoogleAdsBot:  "GoogleAdsBot",
	BrowserLinkedInBot:   "LinkedInBot",
	BrowserMsnBot:        "MsnBot",
	BrowserPingdomBot:    "PingdomBot",
	BrowserTwitterBot:    "TwitterBot",
	BrowserYandexBot:     "YandexBot",
	BrowserCocCocBot:     "CocCocBot",
	BrowserPinterestBot:  "PinterestBot",
	BrowserSlackBot:      "SlackBot",
	BrowserSeekportBot:   "SeekportBot",
	BrowserYahooBot:      "YahooBot",
}

func (b BrowserID) String() string {
	return browserNames[b]
}

// OSID (int) returns a constant.
type OSID uint8

// A complete list of supported OSes in the
// form of constants. For handling particular versions
// of operating systems (e.g. Windows 2000), see
// the README.md file.
const (
	OSUnknown      OSID = 0
	OSWindowsPhone OSID = 1
	OSWindows      OSID = 2
	OSMacOSX       OSID = 3
	OSiOS          OSID = 4
	OSAndroid      OSID = 5
	OSBlackberry   OSID = 6
	OSChromeOS     OSID = 7
	OSKindle       OSID = 8
	OSWebOS        OSID = 9
	OSLinux        OSID = 10
	OSPlaystation  OSID = 11
	OSXbox         OSID = 12
	OSNintendo     OSID = 13
	OSBot          OSID = 14
)

var osNames = map[OSID]string{
	OSUnknown:      "Unknown",
	OSWindowsPhone: "WindowsPhone",
	OSWindows:      "Windows",
	OSMacOSX:       "MacOSX",
	OSiOS:          "iOS",
	OSAndroid:      "Android",
	OSBlackberry:   "Blackberry",
	OSChromeOS:     "ChromeOS",
	OSKindle:       "Kindle",
	OSWebOS:        "WebOS",
	OSLinux:        "Linux",
	OSPlaystation:  "Playstation",
	OSXbox:         "Xbox",
	OSNintendo:     "Nintendo",
	OSBot:          "Bot",
}

func (o OSID) String() string {
	return osNames[o]
}

// Platform (int) returns a constant.
type PlatformID uint8

// A complete list of supported platforms in the
// form of constants. Many OSes report their
// true platform, such as Android OS being Linux
// platform.
const (
	PlatformUnknown      PlatformID = 0
	PlatformWindows      PlatformID = 1
	PlatformMac          PlatformID = 2
	PlatformLinux        PlatformID = 3
	PlatformiPad         PlatformID = 4
	PlatformiPhone       PlatformID = 5
	PlatformiPod         PlatformID = 6
	PlatformBlackberry   PlatformID = 7
	PlatformWindowsPhone PlatformID = 8
	PlatformPlaystation  PlatformID = 9
	PlatformXbox         PlatformID = 10
	PlatformNintendo     PlatformID = 11
	PlatformBot          PlatformID = 12
	PlatformAndroid      PlatformID = 13
)

var platformNames = map[PlatformID]string{
	PlatformUnknown:      "Unknown",
	PlatformWindows:      "Windows",
	PlatformMac:          "Mac",
	PlatformAndroid:      "Android",
	PlatformLinux:        "Linux",
	PlatformiPad:         "iPad",
	PlatformiPhone:       "iPhone",
	PlatformiPod:         "iPod",
	PlatformBlackberry:   "Blackberry",
	PlatformWindowsPhone: "WindowsPhone",
	PlatformPlaystation:  "Playstation",
	PlatformXbox:         "Xbox",
	PlatformNintendo:     "Nintendo",
}

func (b PlatformID) String() string {
	return platformNames[b]
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) Less(c Version) bool {
	if v.Major < c.Major {
		return true
	}

	if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	}

	if v.Minor > c.Minor {
		return false
	}

	return v.Patch < c.Patch
}

type UserAgent struct {
	Browser  Browser
	OS       OS
	DeviceID DeviceID
}

type Browser struct {
	Name    BrowserID
	Version Version
}

type OS struct {
	Platform PlatformID
	Name     OSID
	Version  Version
}

// Reset resets the UserAgent to it's zero value
func (ua *UserAgent) Reset() {
	ua.Browser = Browser{}
	ua.OS = OS{}
	ua.DeviceID = DeviceUnknown
}

// IsBot returns true if the UserAgent represent a bot
func (ua *UserAgent) IsBot() bool {
	if ua.Browser.Name >= BrowserBot && ua.Browser.Name <= BrowserYahooBot {
		return true
	}
	if ua.OS.Name == OSBot {
		return true
	}
	if ua.OS.Platform == PlatformBot {
		return true
	}
	return false
}

// Parse accepts a raw user agent (string) and returns the UserAgent.
func Parse(ua string) *UserAgent {
	dest := new(UserAgent)
	parse(ua, dest)
	return dest
}

// ParseUserAgent is the same as Parse, but populates the supplied UserAgent.
// It is the caller's responsibility to call Reset() on the UserAgent before
// passing it to this function.
func ParseUserAgent(ua string, dest *UserAgent) {
	parse(ua, dest)
}

func parse(ua string, dest *UserAgent) {
	ua = normalise(ua)
	switch {
	case len(ua) == 0:
		dest.OS.Platform = PlatformUnknown
		dest.OS.Name = OSUnknown
		dest.Browser.Name = BrowserUnknown
		dest.DeviceID = DeviceUnknown

	// stop on on first case returning true
	case dest.evalOS(ua):
	case dest.evalBrowserID(ua):
	default:
		dest.evalBrowserVersion(ua)
		dest.evalDevice(ua)
	}
}

// normalise normalises the user supplied agent string so that
// we can more easily parse it.
func normalise(ua string) string {
	if len(ua) <= 1024 {
		var buf [1024]byte
		ascii := copyLower(buf[:len(ua)], ua)
		if !ascii {
			// Fall back for non ascii characters
			return strings.ToLower(ua)
		}
		return string(buf[:len(ua)])
	}
	// Fallback for unusually long strings
	return strings.ToLower(ua)
}

// copyLower copies a lowercase version of s to b. It assumes s contains only single byte characters
// and will panic if b is nil or is not long enough to contain all the bytes from s.
// It returns early with false if any characters were non ascii.
func copyLower(b []byte, s string) bool {
	for j := 0; j < len(s); j++ {
		c := s[j]
		if c > 127 {
			return false
		}

		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}

		b[j] = c
	}
	return true
}

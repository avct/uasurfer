package uasurfer

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	iosVersion           = regexp.MustCompile("(cpu|iphone) os \\d+")
	osxVersionUnderscore = regexp.MustCompile("os x 10_\\d+")
	osxVersionDot        = regexp.MustCompile("os x 10\\.\\d+")
	wpVerLong            = regexp.MustCompile("windows\\sphone\\sos\\s\\d+")
	wpVerShort           = regexp.MustCompile("windows\\sphone\\s\\d+")
	//kindleTest           = regexp.MustCompile("\\sKF[A-Z]{2,4}\\s")
	androidVersion        = regexp.MustCompile("android \\d+")
	amazonFireFingerprint = regexp.MustCompile("\\s(k[a-z]{3,5}|sd\\d{4}ur)\\s") //tablet or phone
	// windowsVersion        = regexp.MustCompile("windows nt (\\d\\d|\\d\\.\\d)")
	digits = regexp.MustCompile("\\d+")
)

// Retrieve the espoused platform and OS from the User-Agent string
func evalSystem(ua string) (Platform, OSName, int) {

	if len(ua) == 0 {
		return PlatformUnknown, OSUnknown, 0
	}

	s := strings.IndexRune(ua, '(')
	e := strings.IndexRune(ua, ')')
	if s > e {
		s = 0
		e = len(ua)
	}
	if e == -1 {
		e = len(ua)
	}

	agentPlatform := ua[s+1 : e]
	specsEnd := strings.Index(agentPlatform, ";")
	var specs string
	if specsEnd != -1 {
		specs = agentPlatform[:specsEnd]
	} else {
		specs = agentPlatform
	}

	//strict OS & version identification
	switch specs {
	case "android":
		return evalLinux(ua, agentPlatform)
	case "bb10", "playbook":
		return PlatformBlackberry, OSBlackberry, 0
	case "x11", "linux":
		return evalLinux(ua, agentPlatform)
	case "ipad", "iphone", "ipod touch", "ipod":
		return evaliOS(specs, agentPlatform)
	case "macintosh":
		return evalMacintosh(ua)
	}

	// Blackberry
	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") {
		return PlatformBlackberry, OSBlackberry, 0
	}

	// Windows Phone
	if strings.Contains(agentPlatform, "windows phone ") {
		return evalWindowsPhone(agentPlatform)
	}

	// Windows, Xbox
	if strings.Contains(ua, "windows ") {
		return evalWindows(ua)
	}

	// Kindle
	if strings.Contains(ua, "kindle/") || amazonFireFingerprint.MatchString(agentPlatform) {
		return PlatformLinux, OSKindle, 0
	}

	// Linux (broader attempt)
	if strings.Contains(ua, "linux") {
		return evalLinux(ua, agentPlatform)
	}

	// WebOS (non-linux flagged)
	if strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos") {
		return PlatformLinux, OSWebOS, 0
	}

	// Nintendo
	if strings.Contains(ua, "nintendo") {
		return PlatformNintendo, OSNintendo, 0
	}

	// Playstation
	if strings.Contains(ua, "playstation") || strings.Contains(ua, "vita") || strings.Contains(ua, "psp") {
		return PlatformPlaystation, OSPlaystation, 0
	}

	// Android
	if strings.Contains(ua, "android") {
		return evalLinux(ua, agentPlatform)
	}

	return PlatformUnknown, OSUnknown, 0 //default

}

// evalLinux returns the `Platform`, `OSName` and `int` of UAs with
// 'linux' listed as their platform.
func evalLinux(ua string, agentPlatform string) (Platform, OSName, int) {

	// Kindle Fire
	if strings.Contains(ua, "kindle") || amazonFireFingerprint.MatchString(agentPlatform) {
		// get the version of Android if available, though we don't call this OSAndroid
		v, _ := findVersionNumber(agentPlatform, "android ")
		return PlatformLinux, OSKindle, v
	}

	// Android, Kindle Fire
	if strings.Contains(ua, "android") || strings.Contains(ua, "googletv") {
		// Android
		if v, ok := findVersionNumber(agentPlatform, "android "); ok {
			return PlatformLinux, OSAndroid, v
		}
		return PlatformLinux, OSAndroid, 0 //default
	}

	// ChromeOS
	if strings.Contains(ua, "cros") {
		return PlatformLinux, OSChromeOS, 0
	}

	// WebOS
	if strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos") {
		return PlatformLinux, OSWebOS, 0
	}

	// Linux, "Linux-like"
	if strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") || strings.Contains(ua, "debian") || strings.Contains(ua, "ubuntu") {
		return PlatformLinux, OSLinux, 0
	}

	return PlatformLinux, OSLinux, 0 //default

}

// evaliOS returns the `Platform`, `OSName` and `int` of UAs with
// 'ipad' or 'iphone' listed as their platform.
func evaliOS(uaPlatform string, agentPlatform string) (Platform, OSName, int) {

	// iPhone
	if uaPlatform == "iphone" {
		return PlatformiPhone, OSiOS, getiOSVersion(agentPlatform)
	}

	// iPad
	if uaPlatform == "ipad" {
		return PlatformiPad, OSiOS, getiOSVersion(agentPlatform)
	}

	// iPod
	if uaPlatform == "ipod touch" || uaPlatform == "ipod" {
		return PlatformiPod, OSiOS, getiOSVersion(agentPlatform)
	}

	return PlatformiPad, OSUnknown, 0 //default
}

func findVersionNumber(s string, m string) (int, bool) {
	if ind := strings.Index(s, m); ind != -1 {
		v := digits.FindString(s[ind+len(m):])
		return strToInt(v), true
	}
	return 0, false
}

func evalWindowsPhone(agentPlatform string) (Platform, OSName, int) {

	if v, ok := findVersionNumber(agentPlatform, "windows phone os "); ok {
		return PlatformWindowsPhone, OSWindowsPhone, v
	}

	if v, ok := findVersionNumber(agentPlatform, "windows phone "); ok {
		return PlatformWindowsPhone, OSWindowsPhone, v
	}

	return PlatformWindowsPhone, OSUnknown, 0 //default
}

func evalWindows(ua string) (Platform, OSName, int) {

	//Xbox -- it reads just like Windows
	if strings.Contains(ua, "xbox") {
		return PlatformXbox, OSXbox, 6
	}

	// No windows version
	if !strings.Contains(ua, "windows ") {
		return PlatformWindows, OSUnknown, 0
	}

	// Windows 10
	if strings.Contains(ua, "windows nt 10") {
		return PlatformWindows, OSWindows, 10
	}

	// Windows 8
	if strings.Contains(ua, "windows nt 6.2") || strings.Contains(ua, "windows nt 6.3") {
		return PlatformWindows, OSWindows, 8
	}

	// Windows 7
	if strings.Contains(ua, "windows nt 6.1") {
		return PlatformWindows, OSWindows, 7
	}

	// Windows Vista
	if strings.Contains(ua, "windows nt 6.0") {
		return PlatformWindows, OSWindows, 6
	}
	// Windows XP
	if strings.Contains(ua, "windows nt 5.1") || strings.Contains(ua, "windows nt 5.2") || strings.Contains(ua, "windows xp") {
		return PlatformWindows, OSWindows, 5
	}

	// Windows 2000
	if strings.Contains(ua, "windows nt 5.0") || strings.Contains(ua, "windows 2000") {
		return PlatformWindows, OSWindows, 4
	}

	return PlatformWindows, OSUnknown, 0 //default

}

func evalMacintosh(uaPlatformGroup string) (Platform, OSName, int) {

	if v, ok := findVersionNumber(uaPlatformGroup, "os x 10_"); ok {
		return PlatformMac, OSMacOSX, v
	}

	if v, ok := findVersionNumber(uaPlatformGroup, "os x 10."); ok {
		return PlatformMac, OSMacOSX, v
	}

	return PlatformMac, OSUnknown, 0 //default

}

// getiOSVersion accepts the platform portion of a UA string and returns
// an `int` of the major version of iOS, or `0` (unknown) on error.
func getiOSVersion(uaPlatformGroup string) int {
	s := iosVersion.FindString(uaPlatformGroup)
	if len(s) == 0 {
		return 0
	}

	// catpure and trim the last 2 characters; convert to int
	i := strToInt(strings.TrimSpace(s[len(s)-2:]))
	return i
}

// strToInt simply accepts a string and returns an `int`,
// with '0' being default.
func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

package uasurfer

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	iosVersion           = regexp.MustCompile("(cpu|iphone) os \\d+(?:_\\d+)*")
	osxVersionUnderscore = regexp.MustCompile("os x 10_\\d+")
	osxVersionDot        = regexp.MustCompile("os x 10\\.\\d+")
	wpVerLong            = regexp.MustCompile("windows\\sphone\\sos\\s\\d+(?:\\.\\d+)*")
	wpVerShort           = regexp.MustCompile("windows\\sphone\\s\\d+(?:\\.\\d+)*")
	//kindleTest           = regexp.MustCompile("\\sKF[A-Z]{2,4}\\s")
	androidVersion        = regexp.MustCompile("android \\d+(?:\\.\\d+)*")
	amazonFireFingerprint = regexp.MustCompile("\\s(k[a-z]{3,5}|sd\\d{4}ur)\\s") //tablet or phone
	// windowsVersion        = regexp.MustCompile("windows nt (\\d\\d|\\d\\.\\d)")
	digits = regexp.MustCompile("\\d+")
	semver = regexp.MustCompile("\\d+(?:[_\\.]\\d+)*")
)

// Retrieve the espoused platform and OS from the User-Agent string
func evalSystem(ua string) (Platform, OSName, Version) {

	if len(ua) == 0 {
		return PlatformUnknown, OSUnknown, Version{}
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
		return PlatformBlackberry, OSBlackberry, Version{}
	case "x11", "linux":
		return evalLinux(ua, agentPlatform)
	case "ipad", "iphone", "ipod touch", "ipod":
		return evaliOS(specs, agentPlatform)
	case "macintosh":
		return evalMacintosh(ua)
	}

	// Blackberry
	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") {
		return PlatformBlackberry, OSBlackberry, Version{}
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
		return PlatformLinux, OSKindle, Version{}
	}

	// Linux (broader attempt)
	if strings.Contains(ua, "linux") {
		return evalLinux(ua, agentPlatform)
	}

	// WebOS (non-linux flagged)
	if strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos") {
		return PlatformLinux, OSWebOS, Version{}
	}

	// Nintendo
	if strings.Contains(ua, "nintendo") {
		return PlatformNintendo, OSNintendo, Version{}
	}

	// Playstation
	if strings.Contains(ua, "playstation") || strings.Contains(ua, "vita") || strings.Contains(ua, "psp") {
		return PlatformPlaystation, OSPlaystation, Version{}
	}

	// Android
	if strings.Contains(ua, "android") {
		return evalLinux(ua, agentPlatform)
	}

	return PlatformUnknown, OSUnknown, Version{} //default

}

// evalLinux returns the `Platform`, `OSName` and `int` of UAs with
// 'linux' listed as their platform.
func evalLinux(ua string, agentPlatform string) (Platform, OSName, Version) {

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
		return PlatformLinux, OSAndroid, Version{} //default
	}

	// ChromeOS
	if strings.Contains(ua, "cros") {
		return PlatformLinux, OSChromeOS, Version{}
	}

	// WebOS
	if strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos") {
		return PlatformLinux, OSWebOS, Version{}
	}

	// Linux, "Linux-like"
	if strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") || strings.Contains(ua, "debian") || strings.Contains(ua, "ubuntu") {
		return PlatformLinux, OSLinux, Version{}
	}

	return PlatformLinux, OSLinux, Version{} //default

}

// evaliOS returns the `Platform`, `OSName` and `int` of UAs with
// 'ipad' or 'iphone' listed as their platform.
func evaliOS(uaPlatform string, agentPlatform string) (Platform, OSName, Version) {

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

	return PlatformiPad, OSUnknown, Version{} //default
}

func findVersionNumber(s string, m string) (Version, bool) {
	if ind := strings.Index(s, m); ind != -1 {
		v := semver.FindString(s[ind+len(m):])
		return strToVersion(v), true
	}
	return Version{}, false
}

func evalWindowsPhone(agentPlatform string) (Platform, OSName, Version) {

	if v, ok := findVersionNumber(agentPlatform, "windows phone os "); ok {
		return PlatformWindowsPhone, OSWindowsPhone, v
	}

	if v, ok := findVersionNumber(agentPlatform, "windows phone "); ok {
		return PlatformWindowsPhone, OSWindowsPhone, v
	}

	return PlatformWindowsPhone, OSUnknown, Version{} //default
}

func evalWindows(ua string) (Platform, OSName, Version) {

	//Xbox -- it reads just like Windows
	if strings.Contains(ua, "xbox") {
		if v, ok := findVersionNumber(ua, "windows nt "); ok {
			return PlatformXbox, OSXbox, v
		}
		return PlatformXbox, OSXbox, Version{6, 0, 0}
	}

	// No windows version
	if !strings.Contains(ua, "windows ") {
		return PlatformWindows, OSUnknown, Version{}
	}

	if strings.Contains(ua, "windows nt ") {
		if v, ok := findVersionNumber(ua, "windows nt "); ok {
			return PlatformWindows, OSWindows, v
		}
	}

	// Windows XP non-semantic version
	if strings.Contains(ua, "windows xp") {
		return PlatformWindows, OSWindows, Version{5, 1, 0}
	}

	return PlatformWindows, OSUnknown, Version{} //default

}

func evalMacintosh(uaPlatformGroup string) (Platform, OSName, Version) {

	if v, ok := findVersionNumber(uaPlatformGroup, "os x 10_"); ok {
		v.Patch, v.Minor, v.Major = v.Minor, v.Major, 10
		return PlatformMac, OSMacOSX, v
	}

	if v, ok := findVersionNumber(uaPlatformGroup, "os x 10."); ok {
		v.Patch, v.Minor, v.Major = v.Minor, v.Major, 10
		return PlatformMac, OSMacOSX, v
	}

	return PlatformMac, OSUnknown, Version{} //default

}

// getiOSVersion accepts the platform portion of a UA string and returns
// an `int` of the major version of iOS, or `0` (unknown) on error.
func getiOSVersion(uaPlatformGroup string) Version {
	s := iosVersion.FindString(uaPlatformGroup)
	if len(s) == 0 {
		return Version{}
	}

	if l := strings.LastIndex(s, " "); l > -1 {
		s = s[l+1:]
	}

	return strToVersion(s)
}

// strToInt simply accepts a string and returns an `int`,
// with '0' being default.
func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// strToVer accepts a string and returns a Version,
// with '0' being default.
func strToVersion(str string) Version {
	ver := Version{}

	var vs []string
	if strings.Index(str, "_") > -1 {
		vs = strings.Split(str, "_")
	} else {
		vs = strings.Split(str, ".")
	}

	// handle zero prefixed version parts
	for i, s := range vs {
		if len(s) > 1 && s[:1] == "0" {
			vs[i] = vs[i][1:]
			vs = append(vs[:i], append([]string{"0"}, vs[i:]...)...)
		}
	}

	l := len(vs)
	if l > 0 {
		ver.Major = strToInt(vs[0])
	}
	if l > 1 {
		ver.Minor = strToInt(vs[1])
	}
	if l > 2 {
		ver.Patch = strToInt(vs[2])
	}

	return ver
}

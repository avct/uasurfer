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
)

// OS type returns a lowercase name (string) of the operating system, along with its major version (int).
// To allow easy of use with math operators, the version numbers for Mac and Win may be slightly unexpected.
// Here are some examples:
//
// 	For Windows XP (Windows NT 5.1), "windows" is the platform, "xp" is the name, and 5 the version.
// 	For OS X 10.5.1, "mac" is the platform, "os x" the name, and 5 the version.
// 	For Android 5.1, "linux" is the platform, "android" is the name, and 5 the version.
//	For iOS 5.1, "iphone" or "ipad" is the platform, "ios" is the name, and 5 the version.
// type OS struct {
// 	Name    OSName
// 	Version int
// }

// Retrieve the espoused platform and OS from the User-Agent string
func evalSystem(ua string) (Platform, OSName, int) {
	var (
		depth             int
		i                 int
		platformUASection []byte //represents portion of ua string that contains platform information
	)

	// more efficient method of getting the spec string than using native regexp package (TODO: need stronger profiling verification)
	for ; i < len(ua); i++ {
		if ua[i] == '(' {
			depth++
		} else if ua[i] == ')' {
			depth--
			// don't be too hungry for parentheticals, the first group should be enough
			if depth == 0 {
				break
			}
		}

		if depth >= 1 {
			platformUASection = append(platformUASection, ua[i])
		}
	}

	agentPlatform := string(platformUASection)
	specs := strings.Split(agentPlatform, ";")
	specs[0] = strings.TrimPrefix(specs[0], "(")

	//strict OS & version identification
	switch specs[0] {
	case "android":
		return evalLinux(ua, agentPlatform)
	case "bb10", "playbook":
		return PlatformBlackberry, OSBlackberry, 0
	case "x11", "linux":
		return evalLinux(ua, agentPlatform)
	case "ipad", "iphone":
		return evaliOS(specs[0], agentPlatform)
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
		v := strings.TrimPrefix(androidVersion.FindString(agentPlatform), "android ")
		i := strToInt(v)

		return PlatformLinux, OSKindle, i
	}

	// Android, Kindle Fire
	if strings.Contains(ua, "android") {

		// Android
		if androidVersion.MatchString(agentPlatform) {
			v := strings.TrimPrefix(androidVersion.FindString(agentPlatform), "android ")
			i := strToInt(v)

			return PlatformLinux, OSAndroid, i
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
	if strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") {
		return PlatformLinux, OSLinux, 0
	}

	return PlatformLinux, OSUnknown, 0 //default

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

	return PlatformiPad, OSUnknown, 0 //default
}

func evalWindowsPhone(agentPlatform string) (Platform, OSName, int) {

	if strings.Contains(agentPlatform, "windows phone os ") {
		v := strings.TrimPrefix(wpVerLong.FindString(agentPlatform), "windows phone os ")
		i := strToInt(v)

		return PlatformWindowsPhone, OSWindowsPhone, i
	}

	if strings.Contains(agentPlatform, "windows phone ") {
		v := strings.TrimPrefix(wpVerShort.FindString(agentPlatform), "windows phone ")
		i := strToInt(v)

		return PlatformWindowsPhone, OSWindowsPhone, i
	}

	return PlatformWindowsPhone, OSUnknown, 0 //default
}

func evalWindows(ua string) (Platform, OSName, int) {

	//Xbox -- it reads just like Windows
	if strings.Contains(ua, "xbox") {
		return PlatformXbox, OSXbox, 6
	}

	// Windows 10
	if strings.Contains(ua, "windows nt 10") {
		return PlatformWindows, OSWindows10, 10
	}

	// Windows 8
	if strings.Contains(ua, "windows nt 6.2") || strings.Contains(ua, "windows nt 6.3") {
		return PlatformWindows, OSWindows8, 6
	}

	// Windows 7
	if strings.Contains(ua, "windows nt 6.1") {
		return PlatformWindows, OSWindows7, 6
	}

	// Windows Vista
	if strings.Contains(ua, "windows nt 6.0") {
		return PlatformWindows, OSWindowsVista, 6
	}
	// Windows XP
	if strings.Contains(ua, "windows nt 5.1") || strings.Contains(ua, "windows nt 5.2") || strings.Contains(ua, "windows xp") {
		return PlatformWindows, OSWindowsXP, 5
	}

	// Windows 2000
	if strings.Contains(ua, "windows nt 5.0") {
		return PlatformWindows, OSWindows2000, 5
	}

	return PlatformWindows, OSUnknown, 0 //default

}

func evalMacintosh(uaPlatformGroup string) (Platform, OSName, int) {

	if strings.Contains(uaPlatformGroup, "os x 10_") {
		if osxVersionUnderscore.MatchString(uaPlatformGroup) {
			v := strings.TrimPrefix(osxVersionUnderscore.FindString(uaPlatformGroup), "os x 10_")
			i := strToInt(v)

			return PlatformMac, OSMacOSX, i
		}
	}

	if strings.Contains(uaPlatformGroup, "os x 10.") {
		if osxVersionDot.MatchString(uaPlatformGroup) {
			v := strings.TrimPrefix(osxVersionDot.FindString(uaPlatformGroup), "os x 10.")
			i := strToInt(v)

			return PlatformMac, OSMacOSX, i
		}
	}

	return PlatformMac, OSUnknown, 0 //default

}

// getiOSVersion accepts the platform portion of a UA string and returns
// an `int` of the major version of iOS, or `0` (unknown) on error.
func getiOSVersion(uaPlatformGroup string) int {
	v := strings.Split(iosVersion.FindString(uaPlatformGroup), " ")[2]
	i := strToInt(v)

	return i
}

// strToInt simply accepts a string and returns an `int`,
// with '0' being default.
func strToInt(str string) int {
	i, _ := strconv.Atoi(str)

	// if !i {
	// 	i = 0
	// }

	return i
}

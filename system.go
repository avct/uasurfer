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

func (u *UserAgent) evalOS(ua string) bool {

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
		u.evalLinux(ua, agentPlatform)

	case "bb10", "playbook":
		u.OS.Platform = PlatformBlackberry
		u.OS.Name = OSBlackberry

	case "x11", "linux":
		u.evalLinux(ua, agentPlatform)

	case "ipad", "iphone", "ipod touch", "ipod":
		u.evaliOS(specs, agentPlatform)

	case "macintosh":
		u.evalMacintosh(ua)

	default:
		switch {
		// Blackberry
		case strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook"):
			u.OS.Platform = PlatformBlackberry
			u.OS.Name = OSBlackberry

		// Windows Phone
		case strings.Contains(agentPlatform, "windows phone "):
			u.evalWindowsPhone(agentPlatform)

		// Windows, Xbox
		case strings.Contains(ua, "windows "):
			u.evalWindows(ua)

		// Kindle
		case strings.Contains(ua, "kindle/") || amazonFireFingerprint.MatchString(agentPlatform):
			u.OS.Platform = PlatformLinux
			u.OS.Name = OSKindle

		// Linux (broader attempt)
		case strings.Contains(ua, "linux"):
			u.evalLinux(ua, agentPlatform)

		// WebOS (non-linux flagged)
		case strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos"):
			u.OS.Platform = PlatformLinux
			u.OS.Name = OSWebOS

		// Nintendo
		case strings.Contains(ua, "nintendo"):
			u.OS.Platform = PlatformNintendo
			u.OS.Name = OSNintendo

		// Playstation
		case strings.Contains(ua, "playstation") || strings.Contains(ua, "vita") || strings.Contains(ua, "psp"):
			u.OS.Platform = PlatformPlaystation
			u.OS.Name = OSPlaystation

		// Android
		case strings.Contains(ua, "android"):
			u.evalLinux(ua, agentPlatform)

		default:
			u.OS.Platform = PlatformUnknown
			u.OS.Name = OSUnknown
		}
	}

	return u.isBot()
}

func (u *UserAgent) isBot() bool {

	if u.OS.Platform == PlatformBot || u.OS.Name == OSBot {
		u.DeviceType = DeviceComputer
		return true
	}

	if u.Browser.Name == BrowserBot {
		u.OS.Platform = PlatformBot
		u.OS.Name = OSBot
		u.DeviceType = DeviceComputer
		return true
	}

	return false
}

// evalLinux returns the `Platform`, `OSName` and Version of UAs with
// 'linux' listed as their platform.
func (u *UserAgent) evalLinux(ua string, agentPlatform string) {

	switch {
	// Kindle Fire
	case strings.Contains(ua, "kindle") || amazonFireFingerprint.MatchString(agentPlatform):
		// get the version of Android if available, though we don't call this OSAndroid
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSKindle
		u.OS.findVersionNumber(agentPlatform, "android ")

	// Android, Kindle Fire
	case strings.Contains(ua, "android") || strings.Contains(ua, "googletv"):
		// Android
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSAndroid
		u.OS.findVersionNumber(agentPlatform, "android ")

	// ChromeOS
	case strings.Contains(ua, "cros"):
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSChromeOS

	// WebOS
	case strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos"):
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSWebOS

	// Linux, "Linux-like"
	case strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") || strings.Contains(ua, "debian") || strings.Contains(ua, "ubuntu"):
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSLinux

	default:
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSLinux
	}
}

// evaliOS returns the `Platform`, `OSName` and Version of UAs with
// 'ipad' or 'iphone' listed as their platform.
func (u *UserAgent) evaliOS(uaPlatform string, agentPlatform string) {

	switch uaPlatform {
	// iPhone
	case "iphone":
		u.OS.Platform = PlatformiPhone
		u.OS.Name = OSiOS
		u.OS.getiOSVersion(agentPlatform)

	// iPad
	case "ipad":
		u.OS.Platform = PlatformiPad
		u.OS.Name = OSiOS
		u.OS.getiOSVersion(agentPlatform)

	// iPod
	case "ipod touch", "ipod":
		u.OS.Platform = PlatformiPod
		u.OS.Name = OSiOS
		u.OS.getiOSVersion(agentPlatform)

	default:
		u.OS.Platform = PlatformiPad
		u.OS.Name = OSUnknown
	}
}

func (u *UserAgent) evalWindowsPhone(agentPlatform string) {
	u.OS.Platform = PlatformWindowsPhone

	if u.OS.findVersionNumber(agentPlatform, "windows phone os ") || u.OS.findVersionNumber(agentPlatform, "windows phone ") {
		u.OS.Name = OSWindowsPhone
	} else {
		u.OS.Name = OSUnknown
	}
}

func (u *UserAgent) evalWindows(ua string) {

	switch {
	//Xbox -- it reads just like Windows
	case strings.Contains(ua, "xbox"):
		u.OS.Platform = PlatformXbox
		u.OS.Name = OSXbox
		if !u.OS.findVersionNumber(ua, "windows nt ") {
			u.OS.Version.Major = 6
			u.OS.Version.Minor = 0
			u.OS.Version.Patch = 0
		}

	// No windows version
	case !strings.Contains(ua, "windows "):
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSUnknown

	case strings.Contains(ua, "windows nt ") && u.OS.findVersionNumber(ua, "windows nt "):
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSWindows

	case strings.Contains(ua, "windows xp"):
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSWindows
		u.OS.Version.Major = 5
		u.OS.Version.Minor = 1
		u.OS.Version.Patch = 0

	default:
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSUnknown

	}
}

func (u *UserAgent) evalMacintosh(uaPlatformGroup string) {

	u.OS.Platform = PlatformMac

	if u.OS.findVersionNumber(uaPlatformGroup, "os x 10_") || u.OS.findVersionNumber(uaPlatformGroup, "os x 10.") {
		u.OS.Name = OSMacOSX
		u.OS.Version.Patch, u.OS.Version.Minor, u.OS.Version.Major = u.OS.Version.Minor, u.OS.Version.Major, 10
	} else {
		u.OS.Name = OSUnknown
	}
}

func (o *OS) findVersionNumber(s string, m string) bool {
	if ind := strings.Index(s, m); ind != -1 {
		o.Version.parse(semver.FindString(s[ind+len(m):]))
		return true
	}
	return false
}

// getiOSVersion accepts the platform portion of a UA string and returns
// a Version.
func (o *OS) getiOSVersion(uaPlatformGroup string) {
	s := iosVersion.FindString(uaPlatformGroup)
	if len(s) == 0 {
		return
	}

	if l := strings.LastIndex(s, " "); l > -1 {
		s = s[l+1:]
	}

	o.Version.parse(s)
}

// strToInt simply accepts a string and returns a `int`,
// with '0' being default.
func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// strToVer accepts a string and returns a Version,
// with {0, 0, 0} being default.
func (v *Version) parse(str string) {

	var vs []string
	if strings.Index(str, "_") > -1 {
		vs = strings.Split(str, "_")
	} else {
		vs = strings.Split(str, ".")
	}

	i := 0
	for _, s := range vs {
		if i > 2 {
			break
		}

		if len(s) == 0 {
			continue
		}

		// handle 0-prefixed versions
		if s[:1] == "0" {
			switch i {
			case 0:
				v.Major = 0

			case 1:
				v.Minor = 0

			case 2:
				v.Patch = 0
			}
			i++

			k := 0
			for k < len(s) {
				if s[k] != '0' {
					break
				}
				k++
			}
			s = s[k:]

			if len(s) == 0 {
				continue
			}
		}

		switch i {
		case 0:
			v.Major, _ = strconv.Atoi(s)

		case 1:
			v.Minor, _ = strconv.Atoi(s)

		case 2:
			v.Patch, _ = strconv.Atoi(s)
		}
		i++
	}
}

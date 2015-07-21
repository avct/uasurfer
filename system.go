package user_agent_surfer

import (
	"regexp"
	"strconv"
	"strings"
)

// OS type returns a lowercase name (string) of the operating system, along with its major version (int).
// To allow easy of use with math operators, the version numbers for Mac and Win may be slightly unexpected.
// Here are some examples:
//
// 	For Windows XP (Windows NT 5.1), "windows" is the platform, "xp" is the name, and 5 the version.
// 	For OS X 10.5.1, "mac" is the platform, "os x" the name, and 5 the version.
// 	For Android 5.1, "linux" is the platform, "android" is the name, and 5 the version.
//	For iOS 5.1, "iphone" or "ipad" is the platform, "ios" is the name, and 5 the version.
type OS struct {
	Name    string
	Version int
}

// Retrieve the espoused platform and OS from the User-Agent string
func (b *BrowserProfile) evalSystem(ua string) (string, string, int) {
	var (
		platform = ""
		os       = ""
		v        = ""
		depth    = 0
		i        = 0
		pgroup   []byte //represents portion of ua string that contains platform information
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
			pgroup = append(pgroup, ua[i])
		}
	}

	pgroup_string := string(pgroup)
	specs := strings.Split(pgroup_string, ";")
	specs[0] = strings.TrimPrefix(specs[0], "(")

	//strict OS & version identification
	switch specs[0] {
	case "bb10", "playbook":
		platform = "blackberry"
		os = "blackberry"
	case "x11", "linux":
		platform = "linux"
	case "ipad", "iphone":

		if specs[0] == "iphone" {
			platform = "iphone"
		} else {
			platform = "ipad"
		}

		iosVersion, _ := regexp.Compile("(cpu|iphone) os \\d+")
		if iosVersion.MatchString(pgroup_string) {
			os = "ios"
			v = strings.Split(iosVersion.FindString(pgroup_string), " ")[2]
		}

	case "macintosh":
		platform = "mac"

		if strings.Contains(pgroup_string, "os x 10_") {
			osxVersion, _ := regexp.Compile("os x 10_\\d+")
			if osxVersion.MatchString(pgroup_string) {
				os = "os x"
				v = strings.TrimPrefix(osxVersion.FindString(pgroup_string), "os x 10_")
			}
		} else if strings.Contains(pgroup_string, "os x 10.") {
			osxVersion, _ := regexp.Compile("os x 10\\.\\d+")
			if osxVersion.MatchString(pgroup_string) {
				os = "os x"
				v = strings.TrimPrefix(osxVersion.FindString(pgroup_string), "os x 10.")
			}
		}
	}

	//fuzzier OS & version identification
	if os == "" {
		//blackberry
		if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") {
			platform = "blackberry"
			os = "blackberry"
			v = "0" // no version number for blackberry
			//windows phone
		} else if strings.Contains(pgroup_string, "windows phone ") {
			platform = "windows phone"
			if strings.Contains(pgroup_string, "windows phone os ") {
				wpVer, _ := regexp.Compile("windows\\sphone\\sos\\s\\d+")
				os = strings.TrimPrefix(wpVer.FindString(pgroup_string), "windows phone os ")
				v = os
			} else {
				wpVer, _ := regexp.Compile("windows\\sphone\\s\\d+")
				os = strings.TrimPrefix(wpVer.FindString(pgroup_string), "windows phone ")
				v = os
			}
			//windows
		} else if strings.Contains(ua, "windows ") {
			//account for xbox looking just like windows, and xbox strings can also show up on Windows Phone
			if strings.Contains(ua, "xbox") {
				platform = "xbox"
				os = "xbox"
				v = "6"
			} else {
				platform = "windows"
				if strings.Contains(ua, "windows nt 10") {
					os = "10"
					v = "10"
				} else if strings.Contains(ua, "windows nt 6.2") || strings.Contains(ua, "windows nt 6.3") {
					os = "8"
					v = "6"
				} else if strings.Contains(ua, "windows nt 6.1") {
					os = "7"
					v = "6"
				} else if strings.Contains(ua, "windows nt 6.0") {
					os = "vista"
					v = "6"
				} else if strings.Contains(ua, "windows nt 5.1") || strings.Contains(ua, "windows nt 5.2") || strings.Contains(ua, "windows xp") {
					os = "xp"
					v = "5"
				} else if strings.Contains(ua, "windows nt 5.0") {
					os = "2000"
					v = "5"
				}
			}
			//WebOS
		} else if strings.Contains(ua, "webos") || strings.Contains(ua, "hpwos") {
			os = "webos"
			platform = "linux"
			v = "0" // Don't bother with OS version for WebOS
		} else if strings.Contains(ua, "kindle/") {
			os = "kindle"
			platform = "kindle"
			v = "0" // Don't bother with OS version for kindle
		} else if strings.Contains(ua, "cros") {
			os = "chromeos"
			platform = "linux"
			v = "0" // Don't bother with OS version for Chrome OS
			// TODO add kindle fire here -- https://developer.amazon.com/appsandservices/solutions/devices/kindle-fire/specifications/04-user-agent-strings
		} else if strings.Contains(ua, "android") {
			// first check if it's a Kindle -- TODO: test if this may be too expensive a method, but kindle != silk

			kindleTest, _ := regexp.Compile("\\sKF[A-Z]{2,4}\\s")
			if kindleTest.MatchString(pgroup_string) {
				platform = "kindle"
			} else {
				platform = "linux"
			}

			os = "android"
			aVersion, _ := regexp.Compile("android \\d+")
			if aVersion.MatchString(pgroup_string) {
				v = strings.TrimPrefix(aVersion.FindString(pgroup_string), "android ")
			}
		} else if strings.Contains(ua, "nintendo") {
			platform = "nintendo"
			os = "nintendo"
			v = "0"
		} else if strings.Contains(ua, "playstation") || strings.Contains(ua, "vita") || strings.Contains(ua, "psp") {
			platform = "playstation"
			os = "playstation"
			if strings.Contains(ua, "playstation 4") {
				v = "4"
			} else if strings.Contains(ua, "playstation 3") {
				v = "3"
			} else if strings.Contains(ua, "playstation 2") {
				v = "2"
			} else {
				v = "0"
			}
		} else {
			v = "0"
		}
	}

	// remaining unidentified platforms
	if platform == "" {
		if strings.Contains(ua, "linux") || strings.Contains(ua, "unix") { // not all these are linux, but they are for our purposes
			platform = "linux"
		} else {
			platform = "unknown"
		}
	}

	// remaining unidentified OSes
	if os == "" {
		if strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") { // not all these are linux, but they are for our purposes
			platform = "linux"
			os = "linux"
		} else {
			os = "unknown"
		}
	}

	version, _ := strconv.Atoi(v)

	return platform, os, version

}

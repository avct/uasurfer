package uasurfer

import (
	"strings"
)

func (u *UserAgent) evalDevice(ua string) {
	switch {
	case u.Browser.Name == BrowserPuffin:
		return

	case u.Browser.Name == BrowserOperaMini:
		if u.OS.Platform == PlatformiPad {
			u.DeviceType = DeviceTablet
			return
		}
		u.DeviceType = DevicePhone

	case u.OS.Platform == PlatformWindows || u.OS.Platform == PlatformMac || u.OS.Name == OSChromeOS:
		if strings.Contains(ua, "wpdesktop") {
			u.DeviceType = DevicePhone
			return
		}
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.DeviceType = DeviceTablet // windows rt, linux haxor tablets
			return
		}
		u.DeviceType = DeviceComputer

	case u.OS.Platform == PlatformOracle:
		if strings.Contains(ua, "sm-t") || strings.Contains(ua, "gt-p") || strings.Contains(ua, "gt-n") || strings.Contains(ua, "ideatab") {
			u.DeviceType = DeviceTablet
			return
		}
		u.DeviceType = DevicePhone

	case u.OS.Platform == PlatformiPad || strings.Contains(ua, "nook") || strings.Contains(ua, "bntv") || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.DeviceType = DeviceTablet

	// long list of smarttv and tv dongle identifiers - above "phone" check to prevent TVs from being detected as phones
	case strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "afta") || strings.Contains(ua, "aftb") || strings.Contains(ua, "afte") || strings.Contains(ua, "afth") || strings.Contains(ua, "aftk") || strings.Contains(ua, "aftm") || strings.Contains(ua, "aftp") || strings.Contains(ua, "afts") || strings.Contains(ua, "aftt") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio") || strings.Contains(ua, "stb") || strings.Contains(ua, "swisscom-ip") || strings.Contains(ua, "tizen") || strings.Contains(ua, "youview"):
		u.DeviceType = DeviceTV

	case u.OS.Platform == PlatformiPhone || u.OS.Platform == PlatformBlackberry || strings.Contains(ua, "phone"):
		u.DeviceType = DevicePhone

	case strings.Contains(ua, " letv"):
		u.DeviceType = DevicePhone

	case u.OS.Name == OSAndroid:
		if strings.Contains(ua, "tablet") || strings.Contains(ua, "nexus 7") || strings.Contains(ua, "nexus 9") || strings.Contains(ua, "nexus 10") || strings.Contains(ua, "xoom") ||
			strings.Contains(ua, "sm-t") || strings.Contains(ua, "; kf") || strings.Contains(ua, "; t1") || strings.Contains(ua, "lenovo tab") ||
			strings.Contains(ua, "sm-p") || strings.Contains(ua, "gt-p") || strings.Contains(ua, "tab") || strings.Contains(ua, "mediapad") || strings.Contains(ua, "lenovo s5") || strings.Contains(ua, "lenovo b6") || strings.Contains(ua, "lenovo b8") || strings.Contains(ua, "a1-8") || strings.Contains(ua, "sgp") || strings.Contains(ua, "p01") || strings.Contains(ua, "p02") || strings.Contains(ua, "gt-n8") || strings.Contains(ua, "gt-n5") || strings.Contains(ua, "sch-i9") || strings.Contains(ua, "sch-i7") || strings.Contains(ua, "p00a") || strings.Contains(ua, "p00i") || strings.Contains(ua, "p008") || strings.Contains(ua, "a33w") || strings.Contains(ua, "lenovo a76") || strings.Contains(ua, "lenovo a55") || strings.Contains(ua, "lenovo a35") || strings.Contains(ua, "lenovo a30") || strings.Contains(ua, "me173x") {
			u.DeviceType = DeviceTablet
			return
		}

		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") {
			u.DeviceType = DevicePhone
			return
		}

		if strings.Contains(ua, "tablet") || strings.Contains(ua, "nexus 7") || strings.Contains(ua, "nexus 9") || strings.Contains(ua, "nexus 10") || strings.Contains(ua, "xoom") ||
			strings.Contains(ua, "sm-t") || strings.Contains(ua, "; kf") || strings.Contains(ua, "; t1") || strings.Contains(ua, "lenovo tab") {
			u.DeviceType = DeviceTablet
			return
		}

		u.DeviceType = DevicePhone // default to phone

	case u.OS.Platform == PlatformPlaystation || u.OS.Platform == PlatformXbox || u.OS.Platform == PlatformNintendo:
		u.DeviceType = DeviceConsole

	case strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v"):
		u.DeviceType = DeviceWearable

	// specifically above "mobile" string check as Kindle Fire tablets report as "mobile"
	case u.Browser.Name == BrowserSilk || u.OS.Name == OSKindle && !strings.Contains(ua, "sd4930ur"):
		u.DeviceType = DeviceTablet

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.DeviceType = DevicePhone

	case u.OS.Name == OSLinux: // linux goes last since it's in so many other device types (tvs, wearables, android-based stuff)
		u.DeviceType = DeviceComputer

	default:
		u.DeviceType = DeviceUnknown
	}
}

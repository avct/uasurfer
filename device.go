package uasurfer

import (
	"strings"
)

func (u *UserAgent) evalDevice(ua string) {
	switch {

	case u.OS.Platform == PlatformWindows || u.OS.Platform == PlatformMac || u.OS.Name == OSChromeOS:
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.DeviceID = DeviceTablet // windows rt, linux haxor tablets
			return
		}
		u.DeviceID = DeviceComputer

	case u.OS.Platform == PlatformiPad || u.OS.Platform == PlatformiPod || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.DeviceID = DeviceTablet

	case u.OS.Platform == PlatformiPhone || u.OS.Platform == PlatformBlackberry || strings.Contains(ua, "phone"):
		u.DeviceID = DevicePhone

	// long list of smarttv and tv dongle identifiers
	case strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "aftt") || strings.Contains(ua, "aftm") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio"):
		u.DeviceID = DeviceTV

	case u.OS.Name == OSAndroid:
		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") {
			u.DeviceID = DevicePhone
			return
		}

		if strings.Contains(ua, "tablet") || strings.Contains(ua, "nexus 7") || strings.Contains(ua, "nexus 9") || strings.Contains(ua, "nexus 10") || strings.Contains(ua, "xoom") ||
			strings.Contains(ua, "sm-t") || strings.Contains(ua, "; kf") || strings.Contains(ua, "; t1") || strings.Contains(ua, "lenovo tab") {
			u.DeviceID = DeviceTablet
			return
		}

		u.DeviceID = DevicePhone // default to phone

	case u.OS.Platform == PlatformPlaystation || u.OS.Platform == PlatformXbox || u.OS.Platform == PlatformNintendo:
		u.DeviceID = DeviceConsole

	case strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v"):
		u.DeviceID = DeviceWearable

	// specifically above "mobile" string check as Kindle Fire tablets report as "mobile"
	case u.Browser.Name == BrowserSilk || u.OS.Name == OSKindle && !strings.Contains(ua, "sd4930ur"):
		u.DeviceID = DeviceTablet

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.DeviceID = DevicePhone

	case u.OS.Name == OSLinux: // linux goes last since it's in so many other device types (tvs, wearables, android-based stuff)
		u.DeviceID = DeviceComputer

	default:
		u.DeviceID = DeviceUnknown
	}
}

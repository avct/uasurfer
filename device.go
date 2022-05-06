package uasurfer

import (
	"strings"
)

var smartTvPatterns = []string{
	"tv",
	"crkey",
	"googletv",
	"aftb",
	"aftt",
	"aftm",
	"adt-",
	"roku",
	"viera",
	"aquos",
	"dtv",
	"appletv",
	"smarttv",
	"tuner",
	"smart-tv",
	"hbbtv",
	"netcast",
	"vizio",
	"x88",
}

func isSmartTv(ua string) bool {
	for _, p := range smartTvPatterns {
		if strings.Contains(ua, p) {
			return true
		}
	}
	return false
}

func (u *UserAgent) evalDevice(ua string) {
	switch {

	case u.OS.Platform == PlatformPlaystation || u.OS.Platform == PlatformXbox || u.OS.Platform == PlatformNintendo || strings.Contains(ua, "nintendo") || strings.Contains(ua, "xbox") || strings.Contains(ua, "playstation"):
		u.DeviceType = DeviceConsole

	case u.OS.Platform == PlatformWindows || u.OS.Platform == PlatformMac || u.OS.Name == OSChromeOS:
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.DeviceType = DeviceTablet // windows rt, linux haxor tablets
			return
		}
		u.DeviceType = DeviceComputer

	case u.OS.Platform == PlatformiPad || u.OS.Platform == PlatformiPod || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.DeviceType = DeviceTablet

	case u.OS.Platform == PlatformiPhone || u.OS.Platform == PlatformBlackberry || strings.Contains(ua, "phone"):
		u.DeviceType = DevicePhone

	// long list of smarttv and tv dongle identifiers
	case isSmartTv(ua):
		u.DeviceType = DeviceTV

	case u.OS.Name == OSAndroid:
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

		// The Nexus Player is a set-top-box / console
		if strings.Contains(ua, "nexus player") {
			u.DeviceType = DeviceConsole
			return
		}

		u.DeviceType = DevicePhone // default to phone

	case strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v"):
		u.DeviceType = DeviceWearable

	// specifically above "mobile" string check as Kindle Fire tablets report as "mobile"
	case u.Browser.Name == BrowserSilk || u.OS.Name == OSKindle && !strings.Contains(ua, "sd4930ur"):
		u.DeviceType = DeviceTablet

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): // anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.DeviceType = DevicePhone

	case u.OS.Name == OSLinux: // linux goes last since it's in so many other device types (tvs, wearables, android-based stuff)

		// https://developers.whatismybrowser.com/useragents/explore/operating_platform/bravia-4k/
		if strings.Contains(ua, "bravia") {
			u.DeviceType = DeviceTV
			return
		}
		u.DeviceType = DeviceComputer

	default:
		u.DeviceType = DeviceUnknown
	}
}

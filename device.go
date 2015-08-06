package uasurfer

import (
	"strings"
)

// Retrieve and/or deduce the espoused device type running the browser. Returns string enum: Computer, Phone, Tablet, Wearable, TV, Console
func evalDevice(ua string, os OSName, platform Platform, browser BrowserName) DeviceType {
	// NOTE: In Go, else clauses are not common. It is more idiomatic to
	//       return early or use a switch statement
	if os == OSMacOSX || platform == PlatformWindows || os == OSChromeOS {
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			return DeviceTablet // windows rt, linux haxor tablets
		}
		return DeviceComputer
	}

	if strings.Contains(ua, "tablet") || platform == PlatformiPad || platform == PlatformKindle || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook") {
		return DeviceTablet
	}

	if strings.Contains(ua, "phone") || platform == PlatformiPhone || platform == PlatformBlackberry {
		return DevicePhone
	}

	// long list of smarttv and tv dongle identifiers
	if strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio") {
		return DeviceTV
	}

	if os == OSAndroid {
		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") && (!strings.Contains(ua, "nexus 7") || !strings.Contains(ua, "nexus 9") || !strings.Contains(ua, "xoom")) {
			return DevicePhone
		}
		return DeviceTablet
	}

	if platform == PlatformPlaystation || platform == PlatformXbox || platform == PlatformNintendo {
		return DeviceConsole
	}

	if strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v") {
		return DeviceWearable
	}

	if browser == BrowserSilk {
		return DeviceTablet
	}

	if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") { //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		return DevicePhone
	}

	if os == OSLinux { // linux goes last since it's in so many other device types (tvs, wearables)
		return DeviceComputer
	}

	return DeviceUnknown
}

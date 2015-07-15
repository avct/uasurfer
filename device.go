package user_agent_surfer

import (
	"strings"
)

// Retrieve and/or deduce the espoused device type running the browser. Returns string enum: Computer, Phone, Tablet, Wearable, TV, Console
func (b *BrowserProfile) evalDevice(ua string) string {
	device := ""

	if b.OS.Name == "os x" || b.Platform == "windows" || b.OS.Name == "chromeos" {
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			device = "tablet" // windows rt, linux haxor tablets
		} else {
			device = "computer"
		}
	} else if strings.Contains(ua, "tablet") || b.Platform == "ipad" || b.Platform == "kindle" || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook") {
		device = "tablet"
	} else if strings.Contains(ua, "phone") || b.Platform == "iphone" || b.Platform == "blackberry" {
		device = "phone"
	} else if b.Platform == "android" {
		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") && (!strings.Contains(ua, "nexus 7") || !strings.Contains(ua, "nexus 9") || !strings.Contains(ua, "xoom")) {
			device = "phone"
		} else {
			device = "tablet"
		}
	} else if b.Platform == "playstation" || b.Platform == "xbox" || b.Platform == "nintendo" {
		device = "console"
	}

	if device == "" {
		// long list of smarttv and tv dongle identifiers
		if strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio") {
			device = "tv"
		} else if strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v") {
			device = "wearable"
		} else if b.Browser.Name == "silk" {
			device = "tablet"
		} else if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") { //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
			device = "phone"
		} else if b.OS.Name == "linux" { // linux goes last since it's in so many other device types (tvs, wearables)
			device = "computer"
		} else {
			device = "unknown"
		}
	}

	return device
}

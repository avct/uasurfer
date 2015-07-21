// Package user_agent_surfer provides fast and reliable abstraction
// of HTTP User-Agent strings. The BrowserProfile struct contains browser name
// (string), browser version (int), platform name (string), os name (string),
// os version (int), device type (string). The philosophy is to identify only
// technology that holds >1% market share, and to avoid expending resources
// and accuracy on guessing at esoteric UA strings.
package user_agent_surfer

import (
	"strings"
)

// The BrowserProfile type contains all the attributes parsed and inferred from the User-Agent string.
type BrowserProfile struct {
	UA         string
	Browser    Browser
	Platform   string
	OS         OS
	DeviceType string
}

func (b *BrowserProfile) initialize() {
	b.UA = ""
	b.Browser.Name = ""
	b.Browser.Version = 0
	b.Platform = ""
	b.OS.Name = ""
	b.OS.Version = 0
	b.DeviceType = ""
}

// New is the core interface, so to speak. Supplying a user-agent string (string) returns a
// completed BrowserProfile type. If there is no match, "unknown" is returned for strings and
// 0 for ints.
func New(ua string) *BrowserProfile {
	ua = strings.ToLower(ua)
	bp := &BrowserProfile{}
	bp.Parse(ua)
	return bp
}

func (b *BrowserProfile) Parse(ua string) {
	b.initialize()
	b.UA = ua

	b.evalBrowser(ua)
	if b.Browser.Name != "bot" {
		b.Platform, b.OS.Name, b.OS.Version = b.evalSystem(ua)
		b.DeviceType = b.evalDevice(ua)
	}
}

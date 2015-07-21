// Package user_agent_surfer provides fast and reliable abstraction
// of HTTP User-Agent strings. The BrowserProfile struct contains browser name
// (string), browser version (int), platform name (string), os name (string),
// os version (int), device type (string). The philosophy is to identify only
// technology that holds >1% market share, and to avoid expending resources
// and accuracy on guessing at esoteric UA strings.
// TODO: Go package names are usually short avoid underscore. Best to rename it to something like useragent
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
	DeviceType DeviceType
}

// TODO: Browser, DeviceType etc will be set to one of a predefined set of values.
//		 Instead of setting them to string values, define constants for each value and set them the constant values.
//       See DeviceType example below and changes to device.go

type DeviceType int

const (
	// Always start with Unknown since it means that an unitialised
	// value will default to Unknown
	DeviceUnknown DeviceType = iota
	DeviceComputer
	DeviceTablet
	DevicePhone
	DeviceConsole
	DeviceWearable
	DeviceTV
)

func (b *BrowserProfile) initialize() {
	b.UA = ""
	b.Browser.Name = ""
	b.Browser.Version = 0
	b.Platform = ""
	b.OS.Name = ""
	b.OS.Version = 0
	b.DeviceType = DeviceUnknown
}

// New is the core interface, so to speak. Supplying a user-agent string (string) returns a
// completed BrowserProfile type. If there is no match, "unknown" is returned for strings and
// 0 for ints.
// TODO: I would remove the New method. Move the lowercase code into Parse
//       and then only the Parse method is needed.
func New(ua string) *BrowserProfile {
	ua = strings.ToLower(ua)
	obj := &BrowserProfile{}
	obj.Parse(ua)
	return obj
}

func (b *BrowserProfile) Parse(ua string) {
	b.initialize()
	ua = strings.ToLower(ua)
	b.UA = ua

	b.Platform, b.OS.Name, b.OS.Version = b.evalSystem(ua)
	b.Browser.Name, b.Browser.Version = b.evalBrowser(ua)
	b.DeviceType = b.evalDevice(ua)
}

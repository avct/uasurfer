// Package uasurfer provides fast and reliable abstraction
// of HTTP User-Agent strings. The BrowserProfile struct contains browser name
// (string), browser version (int), platform name (string), os name (string),
// os version (int), device type (string). The philosophy is to identify only
// technology that holds >1% market share, and to avoid expending resources
// and accuracy on guessing at esoteric UA strings.
// TODO: Go package names are usually short avoid underscore. Best to rename it to something like useragent
package uasurfer

//go:generate stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go

import (
	"strconv"
	"strings"
)

// The BrowserProfile type contains all the attributes parsed and inferred from the User-Agent string.
type BrowserProfile struct {
	UA         string
	Browser    Browser //TODO flatten
	Platform   Platform
	OS         OS //TODO flatten
	DeviceType DeviceType
}

// TODO: Browser, DeviceType etc will be set to one of a predefined set of values.
//		 Instead of setting them to string values, define constants for each value and set them the constant values.
//       See DeviceType example below and changes to device.go

// DeviceType (int) returns a constant.
type DeviceType int

const (
	DeviceUnknown DeviceType = iota
	DeviceComputer
	DeviceTablet
	DevicePhone
	DeviceConsole
	DeviceWearable
	DeviceTV
)

type BrowserName int

const (
	BrowserUnknown BrowserName = iota
	BrowserChrome
	BrowserIE
	BrowserSafari
	BrowserFirefox
	BrowserAndroid
	BrowserOpera
	BrowserBlackberry
	BrowserUCBrowser
	BrowserSilk
	BrowserNokia
	BrowserBot
)

type OSName int

const (
	OSUnknown OSName = iota
	OSWindowsPhone
	OSWindows2000
	OSWindowsXP
	OSWindowsVista
	OSWindows7
	OSWindows8
	OSWindows10
	OSMacOSX
	OSiOS
	OSAndroid
	OSBlackberry
	OSChromeOS
	OSKindle
	OSWebOS
	OSLinux
	OSPlaystation
	OSXbox
	OSNintendo
	OSBot
)

type Platform int

const (
	PlatformUnknown Platform = iota
	PlatformWindows
	PlatformMac
	PlatformLinux
	PlatformiPad
	PlatformiPhone
	PlatformBlackberry
	PlatformWindowsPhone
	PlatformKindle
	PlatformPlaystation
	PlatformXbox
	PlatformNintendo
	PlatformBot
)

func (b *BrowserProfile) initialize() {
	b.UA = ""
	b.Browser.Name = BrowserUnknown
	b.Browser.Version = 0
	b.Platform = PlatformUnknown
	b.OS.Name = OSUnknown
	b.OS.Version = 0
	b.DeviceType = DeviceUnknown
}

func (b *BrowserProfile) Parse(ua string) {
	b.initialize()
	ua = strings.ToLower(ua)

	b.UA = ua
	b.Platform, b.OS.Name, b.OS.Version = b.evalSystem(ua)
	b.Browser.Name = b.evalBrowserName(ua)
	b.Browser.Version, _ = strconv.Atoi(b.evalBrowserVersion(ua))
	b.DeviceType = evalDevice(ua, b.OS.Name, b.Platform, b.Browser.Name)
}

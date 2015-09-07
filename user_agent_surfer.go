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
	"strings"
)

// The BrowserProfile type contains all the attributes parsed and inferred from the User-Agent string.
// type BrowserProfile struct {
// 	UA             string
// 	Browser        BrowserName
// 	BrowserVersion int
// 	Platform       Platform
// 	OS             OSName
// 	OSVersion      int
// 	DeviceType     DeviceType
// }

// DeviceType (int) returns a constant.
type DeviceType int

const (
	DeviceUnknown DeviceType = iota
	Computer
	Tablet
	Phone
	Console
	Wearable
	TV
)

// BrowserName (int) returns a constant.
type BrowserName int

const (
	BrowserUnknown BrowserName = iota
	Chrome
	IE
	Safari
	Firefox
	Android
	Opera
	Blackberry
	UCBrowser
	Silk
	Nokia
	Spotify
	Bot
)

// OSName (int) returns a constant.
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

// Platform (int) returns a constant.
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
	PlatformPlaystation
	PlatformXbox
	PlatformNintendo
	PlatformBot
)

// func (b *BrowserProfile) initialize() {
// 	b.UA = ""
// 	b.Browser.Name = BrowserUnknown
// 	b.Browser.Version = 0
// 	b.Platform = PlatformUnknown
// 	b.OS.Name = OSUnknown
// 	b.OS.Version = 0
// 	b.DeviceType = DeviceUnknown
// }

// Parse accepts a raw user agent (string) and returns the
// browser name (int), browser version
// (int), platform (int), OS name (int), OS version (int),
// device type (int), and raw user agent (string).
func Parse(ua string) (BrowserName, int, Platform, OSName, int, DeviceType, string) {
	ua = strings.ToLower(ua)

	platform, osName, osVersion := evalSystem(ua)
	browserName := evalBrowserName(ua)
	browserVersion := evalBrowserVersion(ua, browserName)
	deviceType := evalDevice(ua, osName, platform, browserName)

	return browserName, browserVersion, platform, osName, osVersion, deviceType, ua
}

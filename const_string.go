// generated by stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go; DO NOT EDIT

package uasurfer

import "fmt"

const _DeviceType_name = "DeviceUnknownComputerTabletPhoneConsoleWearableTV"

var _DeviceType_index = [...]uint8{0, 13, 21, 27, 32, 39, 47, 49}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceType_index)-1) {
		return fmt.Sprintf("DeviceType(%d)", i)
	}
	return _DeviceType_name[_DeviceType_index[i]:_DeviceType_index[i+1]]
}

const _BrowserName_name = "BrowserUnknownChromeIESafariFirefoxAndroidOperaBlackberryUCBrowserSilkNokiaSpotifyBot"

var _BrowserName_index = [...]uint8{0, 14, 20, 22, 28, 35, 42, 47, 57, 66, 70, 75, 82, 85}

func (i BrowserName) String() string {
	if i < 0 || i >= BrowserName(len(_BrowserName_index)-1) {
		return fmt.Sprintf("BrowserName(%d)", i)
	}
	return _BrowserName_name[_BrowserName_index[i]:_BrowserName_index[i+1]]
}

const _OSName_name = "OSUnknownOSWindowsPhoneOSWindows2000OSWindowsXPOSWindowsVistaOSWindows7OSWindows8OSWindows10OSMacOSXOSiOSOSAndroidOSBlackberryOSChromeOSOSKindleOSWebOSOSLinuxOSPlaystationOSXboxOSNintendoOSBot"

var _OSName_index = [...]uint8{0, 9, 23, 36, 47, 61, 71, 81, 92, 100, 105, 114, 126, 136, 144, 151, 158, 171, 177, 187, 192}

func (i OSName) String() string {
	if i < 0 || i >= OSName(len(_OSName_index)-1) {
		return fmt.Sprintf("OSName(%d)", i)
	}
	return _OSName_name[_OSName_index[i]:_OSName_index[i+1]]
}

const _Platform_name = "PlatformUnknownPlatformWindowsPlatformMacPlatformLinuxPlatformiPadPlatformiPhonePlatformBlackberryPlatformWindowsPhonePlatformPlaystationPlatformXboxPlatformNintendoPlatformBot"

var _Platform_index = [...]uint8{0, 15, 30, 41, 54, 66, 80, 98, 118, 137, 149, 165, 176}

func (i Platform) String() string {
	if i < 0 || i >= Platform(len(_Platform_index)-1) {
		return fmt.Sprintf("Platform(%d)", i)
	}
	return _Platform_name[_Platform_index[i]:_Platform_index[i+1]]
}

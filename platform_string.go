// Code generated by "stringer -type=Platform -trimprefix=Platform"; DO NOT EDIT.

package uasurfer

import "strconv"

const _Platform_name = "UnknownWindowsMacLinuxiPadiPhoneiPodBlackberryWindowsPhonePlaystationXboxNintendoBot"

var _Platform_index = [...]uint8{0, 7, 14, 17, 22, 26, 32, 36, 46, 58, 69, 73, 81, 84}

func (i Platform) String() string {
	if i < 0 || i >= Platform(len(_Platform_index)-1) {
		return "Platform(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Platform_name[_Platform_index[i]:_Platform_index[i+1]]
}
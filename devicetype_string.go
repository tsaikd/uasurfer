// Code generated by "stringer -type=DeviceType -trimprefix=Device"; DO NOT EDIT.

package uasurfer

import "strconv"

const _DeviceType_name = "UnknownComputerTabletPhoneConsoleWearableTV"

var _DeviceType_index = [...]uint8{0, 7, 15, 21, 26, 33, 41, 43}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceType_index)-1) {
		return "DeviceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceType_name[_DeviceType_index[i]:_DeviceType_index[i+1]]
}

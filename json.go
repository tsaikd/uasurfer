package uasurfer

import "encoding/json"

func (t DeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t BrowserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t OSName) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

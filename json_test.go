package uasurfer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_json(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	type testStruct struct {
		Device   DeviceType  `json:",omitempty"`
		Browser  BrowserName `json:",omitempty"`
		OS       OSName      `json:",omitempty"`
		Platform Platform    `json:",omitempty"`
	}

	if data, err := json.Marshal(testStruct{Device: DeviceComputer}); assert.NoError(err) {
		require.EqualValues(`{"Device":"Computer"}`, string(data))
	}
	if data, err := json.Marshal(testStruct{Browser: BrowserChrome}); assert.NoError(err) {
		require.EqualValues(`{"Browser":"Chrome"}`, string(data))
	}
	if data, err := json.Marshal(testStruct{OS: OSMacOSX}); assert.NoError(err) {
		require.EqualValues(`{"OS":"MacOSX"}`, string(data))
	}
	if data, err := json.Marshal(testStruct{Platform: PlatformMac}); assert.NoError(err) {
		require.EqualValues(`{"Platform":"Mac"}`, string(data))
	}
}

package loaders

import (
	"io/ioutil"
	"testing"
)

func TestLoadPlist(t *testing.T) {
	tests := []string{
		"plist/testdata/file.plist",
		// "plist/testdata/C.plist",
	}

	var v struct{}
	for i, test := range tests {
		if d, err := ioutil.ReadFile(test); err != nil {
			t.Errorf("Test %d: Unable to read %s: %s", i, test, err)
		} else if err := LoadPlist(d, &v); err != nil {
			t.Errorf("Test %d: Unable to load plist %s: %s", i, test, err)
		}
	}
}

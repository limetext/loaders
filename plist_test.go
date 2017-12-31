package loaders

import (
	"io/ioutil"
	"testing"
)

func TestLoadPlist(t *testing.T) {
	var v struct{}
	if d, err := ioutil.ReadFile("testdata/file.plist"); err != nil {
		t.Errorf("Unable to read plist file: %s", err)
	} else if err := LoadPlist(d, &v); err != nil {
		t.Errorf("Unable to load plist file: %s", err)
	}
}

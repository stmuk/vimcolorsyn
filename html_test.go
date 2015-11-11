package vimcolorsyn

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {

	fn := "/tmp/vimcolorsyn_test.html"

	os.Remove(fn)

	v := &Vimcolour{"vimcolorsyn.go", fn, "go"}
	v.New()

	f, err := ioutil.ReadFile(fn)

	if err != nil {
		t.Errorf("filename %s doesn't exist", fn)
	}

	ceRegex, err := regexp.Compile(`vimCodeElement`)

	if !ceRegex.Match(f) {
		t.Errorf("vimCodeElement not found")
	}

	pRegex, err := regexp.Compile(`<span class="Statement">package</span> vimcolorsyn`)
	if !pRegex.Match(f) {
		t.Errorf("package vimcolorsyn fragment not found")
	}

}

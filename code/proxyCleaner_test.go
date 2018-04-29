package code

import (
	"testing"
)

func TestProxyCleaner(t *testing.T) {

	link := Unique(GetTag())

	if с := Cleaner(link[0]); с == "" {
		t.Fail()
	}
}

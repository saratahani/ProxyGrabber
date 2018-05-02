package code

import (
	"testing"
)

func TestProxyCleaner(t *testing.T) {

	link := unique(getTag())

	if с := cleaner(link[0]); с == "" {
		t.Fail()
	}
}

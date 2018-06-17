package grabber

import (
	"testing"
)

func TestFetchFreshProxies(t *testing.T) {
	if err := FetchFreshProxies(); err != nil {
		t.Fail()
	}
}

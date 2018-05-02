package code

import (
	"testing"
)

func TestCheckProxySOCKS(t *testing.T) {

	ch := make(chan QR, 1)

	checkProxySOCKS(`000.00.000.000:0000`, ch)

	r := <-ch

	if r.Res != false {
		t.Fail()
	}

	checkProxySOCKS(`51.15.100.63:1080`, ch)

	r = <-ch

	if r.Res != true {
		t.Fail()
	}
}

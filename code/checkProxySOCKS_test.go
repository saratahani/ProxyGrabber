package code

import (
	"sync"
	"testing"
)

func TestCheckProxySOCKS(t *testing.T) {

	var wg sync.WaitGroup

	ch := make(chan QR, 1)

	wg.Add(1)
	checkProxySOCKS(`000.00.000.000:0000`, ch, &wg)

	r := <-ch

	wg.Wait()

	if r.Res != false {
		t.Fail()
	}

	wg.Add(1)
<<<<<<< HEAD
	checkProxySOCKS(`139.162.182.186:58482`, ch, &wg)
=======
	checkProxySOCKS(`91.106.187.254:39880`, ch, &wg)
>>>>>>> b665c3f1a5d3232fbff6e5aa7d496646cb2258d3

	r = <-ch

	wg.Wait()

	if r.Res != true {
		t.Fail()
	}

	close(ch)
}

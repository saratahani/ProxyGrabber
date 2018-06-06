package code

import (
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	timeout       = time.Duration(1500 * time.Millisecond)
	dialerTimeout = time.Duration(200 * time.Millisecond)
)

// checkProxySOCKS Check proxies on valid
func checkProxySOCKS(prox string, c chan QR, wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	d := net.Dialer{
		Timeout:   dialerTimeout,
		KeepAlive: dialerTimeout,
	}

	dialer, _ := proxy.SOCKS5("tcp", prox, nil, &d)

	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial:              dialer.Dial,
		},
	}

	res, err := httpClient.Get("https://telegram.org/")
	if err != nil {
		c <- QR{Addr: prox, Res: false}
		return
	}

	defer res.Body.Close()
	io.Copy(ioutil.Discard, res.Body)

	c <- QR{Addr: prox, Res: true}

	return nil
}

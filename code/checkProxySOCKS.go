package code

import (
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//checkProxySOCKS Check proxies on valid
func checkProxySOCKS(prox string, c chan QR, wg *sync.WaitGroup) (err error) {

	defer wg.Done()

	//Sending request through proxy
	dialer, _ := proxy.SOCKS5("tcp", prox, nil, proxy.Direct)

	timeout := time.Duration(2 * time.Second)

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

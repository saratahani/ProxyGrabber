package code

import (
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//CheckProxySOCKS Check proxies on valid
func CheckProxySOCKS(prox string, c chan QR) (err error) {

	//Sending request through proxy
	dialer, _ := proxy.SOCKS5("tcp", prox, nil, proxy.Direct)
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout, Transport: &http.Transport{DisableKeepAlives: true, Dial: dialer.Dial}}
	res, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {

		c <- QR{Addr: prox, Res: false}
		return
	}

	_, err = io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	c <- QR{Addr: prox, Res: true}
	return
}

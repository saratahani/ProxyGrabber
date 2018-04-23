package code

import (
	"golang.org/x/net/proxy"
	"net/http"
	"time"
)

//CheckProxySOCKS Check proxies on valid
func CheckProxySOCKS(proxyy string, c chan QR) {

	//Sending request through proxy
	dialer, _ := proxy.SOCKS5("tcp", proxyy, nil, proxy.Direct)
	timeout := time.Duration(20 * time.Second)
	httpClient := &http.Client{Timeout: timeout, Transport: &http.Transport{Dial: dialer.Dial}}
	_, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {

		c <- QR{Addr: proxyy, Res: false}
	} else {

		c <- QR{Addr: proxyy, Res: true}
	}
}

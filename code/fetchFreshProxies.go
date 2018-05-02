package code

import (
	"strings"
)

//UniqueProxies contain field Proxy
type UniqueProxies struct {
	Proxy []string
}

//UP contains unique checked proxies
var UP UniqueProxies

var linksArray, checkedProxiesArray, splitedProxies []string

//FetchFreshProxies fetching and checking proxies
func FetchFreshProxies() error {

	respChan := make(chan QR)

	//creating array with links
	for _, v := range unique(getTag()) {
		cleanLinks := cleaner(v)
		linksArray = append(linksArray, cleanLinks)
	}

	//splitting proxies
	for _, val := range linksArray {
		splitedProxies = append(splitedProxies, strings.Split(val, "\n")...)
	}

	//checking proxies
	for _, proxy := range splitedProxies {
		go checkProxySOCKS(proxy, respChan)
	}

	for range splitedProxies {
		r := <-respChan
		if r.Res {
			checkedProxiesArray = append(checkedProxiesArray, r.Addr)
		}
	}

	//checking proxies on uniqueness
	UP.Proxy = unique(checkedProxiesArray)

	//reset arrays
	linksArray = nil
	splitedProxies = nil
	checkedProxiesArray = nil

	return nil
}

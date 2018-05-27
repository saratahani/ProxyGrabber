package code

import (
	"strings"
	"sync"
)

// UniqueProxies contain field Proxy
type UniqueProxies struct {
	Proxy []string
}

// UP contains unique checked proxies
var UP UniqueProxies

var linksArray, checkedProxiesArray, splitedProxies []string

var wg sync.WaitGroup

// FetchFreshProxies fetching and checking proxies
func FetchFreshProxies() error {

	respChan := make(chan QR)

	// creating array with links
	for _, v := range unique(getTag()) {
		cleanLinks := cleaner(v)
		linksArray = append(linksArray, cleanLinks)
	}

	// splitting proxies
	for _, val := range linksArray {
		splitedProxies = append(splitedProxies, strings.Split(val, "\n")...)
	}

	// checking proxies
	for _, proxy := range splitedProxies {
		wg.Add(1)
		go checkProxySOCKS(proxy, respChan, &wg)
	}

	for range splitedProxies {
		wg.Add(1)
		r := <-respChan
		if r.Res {
			checkedProxiesArray = append(checkedProxiesArray, r.Addr)
		}
		wg.Done()
	}

	wg.Wait()

	// checking proxies on uniqueness
	UP.Proxy = unique(checkedProxiesArray)

	// reset arrays
	linksArray = nil
	splitedProxies = nil
	checkedProxiesArray = nil

	close(respChan)

	return nil
}

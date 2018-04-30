package main

import (
	"encoding/json"
	"github.com/trigun117/ProxyGrabber/code"
	"net/http"
	"strings"
	"time"
)

var linksArray, checkedProxiesArray, splitedProxies, uniqueProxies []string

func fetchFreshProxies() {

	//reset arrays
	linksArray = nil
	splitedProxies = nil
	checkedProxiesArray = nil

	respChan := make(chan code.QR)

	//creating array with links
	for _, v := range code.Unique(code.GetTag()) {
		cleanLinks := code.Cleaner(v)
		linksArray = append(linksArray, cleanLinks)
	}

	//splitting proxies
	for _, val := range linksArray {
		splitedProxies = append(splitedProxies, strings.Split(val, "\n")...)
	}

	//checking proxies
	for _, proxy := range splitedProxies {
		go code.CheckProxySOCKS(proxy, respChan)
	}

	for range splitedProxies {
		r := <-respChan
		if r.Res {
			checkedProxiesArray = append(checkedProxiesArray, r.Addr)
		}
	}

	//checking proxies on uniqueness
	uniqueProxies = code.Unique(checkedProxiesArray)
}

//browser cache
func cacheHandler(h http.Handler, n string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age="+n)
		h.ServeHTTP(w, r)
	})
}

//json response
func sendJSONHandler(w http.ResponseWriter, r *http.Request) {

	j := struct {
		Proxies []string
	}{Proxies: uniqueProxies}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(j)
}

func server() {

	go func() {
		for {
			fetchFreshProxies()
			time.Sleep(2 * time.Minute)
		}

	}()

	http.Handle("/", cacheHandler(http.FileServer(http.Dir("./template/index")), "900"))

	http.HandleFunc("/json", sendJSONHandler)

	//loading template files
	http.Handle("/static/", http.StripPrefix("/static/", cacheHandler(http.FileServer(http.Dir("./template/static")), "31536000")))

	http.ListenAndServe(":80", nil)

}

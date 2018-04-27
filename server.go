package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trigun117/ProxyGrabber/code"
	"net/http"
	"strings"
	"time"
)

var linksArray, checkedProxiesArray, splitedProxies, uniqueProxies []string

func fetchFreshProxies() {
	respChan := make(chan code.QR)

	//creating array with links
	for _, v := range code.Unique(getTag()) {
		cleanLinks, _ := cleaner(v)
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

func sendJSON(w http.ResponseWriter, r *http.Request) {

	j := struct {
		Proxies []string
	}{Proxies: uniqueProxies}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(j)
}

func server() {

	router := mux.NewRouter()

	go func() {
		for {
			fetchFreshProxies()
			time.Sleep(5 * time.Minute)
		}

	}()

	router.HandleFunc("/json", sendJSON)

	//loading template files
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":80", router)

}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var joinedProxies string
var linksArray, checkedProxiesArray, splitedProxies, uniqueProxies []string

func fetchFreshProxies() {
	respChan := make(chan QR)

	for _, v := range unique(getTag()) {
		cleanLinks, _ := cleaner(v)
		linksArray = append(linksArray, cleanLinks)
	}

	for _, val := range linksArray {
		splitedProxies = append(splitedProxies, strings.Split(val, "\n")...)
	}

	for _, proxy := range splitedProxies {
		go checkProxySOCKS(proxy, respChan)
	}

	for range splitedProxies {
		r := <-respChan
		if r.Res {
			checkedProxiesArray = append(checkedProxiesArray, r.Addr)
		}
	}

	uniqueProxies = unique(checkedProxiesArray)

	joinedProxies = strings.Join(uniqueProxies, "\n")
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	randomNumber := random(0, len(uniqueProxies))
	randomProxy := strings.Split(uniqueProxies[randomNumber], ":")
	link := `tg://socks?server=` + randomProxy[0] + `&port=` + randomProxy[1]
	setProxy := template.URL(link)

	t, _ := template.ParseFiles("./template/index.html")

	p := struct {
		Proxies string
		Link    template.URL
	}{Proxies: joinedProxies, Link: setProxy}

	t.Execute(w, p)
}

func sendJSON(w http.ResponseWriter, r *http.Request) {

	j := struct {
		Proxies []string
	}{Proxies: uniqueProxies}

	json.NewEncoder(w).Encode(j)
}

func server() {

	router := mux.NewRouter()

	go func() {
		for {
			println("FETCHING..")
			fetchFreshProxies()
			time.Sleep(10 * time.Minute)
		}

	}()

	time.Sleep(30 * time.Second)

	println("Listening..")

	router.PathPrefix("/template/").Handler(http.StripPrefix("/template/", http.FileServer(http.Dir("./template/"))))

	router.HandleFunc("/json", sendJSON)

	router.HandleFunc("/", mainPage)

	http.ListenAndServe(":80", router)

}

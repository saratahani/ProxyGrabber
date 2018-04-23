package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trigun117/ProxyGrabber/code"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var joinedProxies string
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

	//joining proxies into one string
	joinedProxies = strings.Join(uniqueProxies, "\n")
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	randomNumber := code.Random(0, len(uniqueProxies))
	randomProxy := strings.Split(uniqueProxies[randomNumber], ":")
	link := `tg://socks?server=` + randomProxy[0] + `&port=` + randomProxy[1]
	setProxy := template.URL(link)

	t, _ := template.ParseFiles("./code/template/index.html")

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

	//loading template files
	router.PathPrefix("/code/template/").Handler(http.StripPrefix("/code/template/", http.FileServer(http.Dir("./code/template/"))))

	router.HandleFunc("/json", sendJSON)

	router.HandleFunc("/", mainPage)

	http.ListenAndServe(":80", router)

}

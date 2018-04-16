package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var l string
var s, m, q, d []string

func fetchFreshProxies() {
	respChan := make(chan QR)

	for _, v := range unique(getTag()) {
		r, _ := cleaner(v)
		s = append(s, r)
	}

	for _, val := range s {
		q = append(q, strings.Split(val, "\n")...)
	}

	for _, proxy := range q {
		go checkProxySOCKS(proxy, respChan)
	}

	for range q {
		r := <-respChan
		if r.Res {
			m = append(m, r.Addr)
		}
	}

	d = unique(m)

	l = strings.Join(d, "\n")
}

func pr(w http.ResponseWriter, r *http.Request) {

	v := random(0, len(d))
	c := strings.Split(d[v], ":")
	link := `tg://socks?server=` + c[0] + `&port=` + c[1]
	i := template.URL(link)

	t, _ := template.ParseFiles("index.html")

	p := struct {
		Proxies string
		Link    template.URL
	}{Proxies: l, Link: i}

	t.Execute(w, p)
}

func sendJSON(w http.ResponseWriter, r *http.Request) {

	j := struct {
		Proxies []string
	}{Proxies: d}

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

	time.Sleep(40 * time.Second)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/json", sendJSON)

	router.HandleFunc("/", pr)

	http.ListenAndServe(":80", router)

}

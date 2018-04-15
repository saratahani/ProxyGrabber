package main

import (
	"html/template"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var l string

func fetchFreshProxies() {

	var s, m, q []string

	runtime.GOMAXPROCS(runtime.NumCPU())

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

	d := unique(m)

	l = strings.Join(d, "\n")
}

func pr(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	p := struct {
		Proxies string
	}{Proxies: l}
	t.Execute(w, p)
}

func server() {

	go func() {
		for {
			println("FETCHING..")
			fetchFreshProxies()
			time.Sleep(10 * time.Minute)
		}

	}()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.HandleFunc("/", pr)

	http.ListenAndServe(":80", nil)

}

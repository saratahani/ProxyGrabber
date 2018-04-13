package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

var l string

func pr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", l)
}

func main() {
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

	http.HandleFunc("/", pr)

	http.ListenAndServe(":80", nil)

}

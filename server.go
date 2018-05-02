package main

import (
	"encoding/json"
	"github.com/trigun117/ProxyGrabber/code"
	"net/http"
	"runtime"
	"time"
)

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
	}{Proxies: code.UP.Proxy}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(j)
}

func server() {

	go func() {
		for {
			code.FetchFreshProxies()
			runtime.GC()
			time.Sleep(2 * time.Minute)
		}

	}()

	http.Handle("/", cacheHandler(http.FileServer(http.Dir("./template/index")), "900"))

	http.HandleFunc("/json", sendJSONHandler)

	//loading template files
	http.Handle("/static/", http.StripPrefix("/static/", cacheHandler(http.FileServer(http.Dir("./template/static")), "31536000")))

	http.ListenAndServe(":80", nil)

}

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/trigun117/ProxyGrabber/grabber"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"runtime"
	"strings"
	"time"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

var (
	emailFrom         = os.Getenv("EF")
	emailTo           = os.Getenv("ET")
	emailFromLogin    = os.Getenv("EFL")
	emailFromPassword = os.Getenv("EFP")
	apiPassword       = os.Getenv("APIPAS")
	corsAddrSite      = os.Getenv("CORSS")
)

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// gzip compression
func gzipHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set cache control
		w.Header().Set("Cache-Control", "max-age=600")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		h.ServeHTTP(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

// json response
func sendJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "template/api/api.html")
	} else if r.Method == "POST" {
		r.ParseForm()
		if r.Form["password"][0] == apiPassword {
			j := struct {
				Proxies []string
			}{Proxies: grabber.UP.Proxy}
			w.Header().Set("Access-Control-Allow-Origin", corsAddrSite)
			json.NewEncoder(w).Encode(j)
		} else {
			http.ServeFile(w, r, "template/api/api.html")
		}
	}
}

// contact form
func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "template/contact/contact.html")
	} else if r.Method == "POST" {
		r.ParseForm()
		body := fmt.Sprintf("Name: %s\n Email: %s\n Message: %s", r.Form["name"][0], r.Form["email"][0], r.Form["message"][0])
		msg := fmt.Sprintf("From: %s \nTo: %s \nSubject: Contact\n\n %s", emailFrom, emailTo, body)
		smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", emailFromLogin, emailFromPassword, "smtp.gmail.com"), emailFrom, []string{emailTo}, []byte(msg))
		http.Redirect(w, r, "/", 301)
	}
}

func server() {
	go func() {
		for {
			defer runtime.GC()
			grabber.FetchFreshProxies()
			time.Sleep(2 * time.Minute)
		}
	}()

	http.Handle("/", http.FileServer(http.Dir("./template/index")))
	http.HandleFunc("/json", sendJSONHandler)
	http.HandleFunc("/contact", contactHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./template/static"))))
	http.ListenAndServe(":80", gzipHandler(http.DefaultServeMux))
}

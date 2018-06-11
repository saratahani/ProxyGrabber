package main

import (
	"encoding/json"
	"fmt"
	"github.com/trigun117/ProxyGrabber/code"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"runtime"
	"time"
)

var (
	emailFrom         = os.Getenv("EF")
	emailTo           = os.Getenv("ET")
	emailFromLogin    = os.Getenv("EFL")
	emailFromPassword = os.Getenv("EFP")
	apiPassword       = os.Getenv("APIPAS")
	corsAddrSite      = os.Getenv("CORSS")
)

// browser cache
func cacheHandler(h http.Handler, n string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age="+n)
		h.ServeHTTP(w, r)
	})
}

func apiTemplate() (t *template.Template) {
	t, _ = template.ParseFiles("template/api/api.html")
	return
}

// json response
func sendJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		apiTemplate().Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		if r.Form["password"][0] == apiPassword {
			j := struct {
				Proxies []string
			}{Proxies: code.UP.Proxy}
			w.Header().Set("Access-Control-Allow-Origin", corsAddrSite)
			json.NewEncoder(w).Encode(j)
		} else {
			fmt.Fprintln(w, "<script>alert('Wrong Password')</script>")
			apiTemplate().Execute(w, nil)
		}
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/contact/contact.html")
		t.Execute(w, nil)
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
			code.FetchFreshProxies()
			time.Sleep(2 * time.Minute)
		}
	}()

	http.Handle("/", cacheHandler(http.FileServer(http.Dir("./template/index")), "600"))
	http.HandleFunc("/json", sendJSONHandler)
	http.HandleFunc("/contact", contactHandler)
	http.Handle("/static/", http.StripPrefix("/static/", cacheHandler(http.FileServer(http.Dir("./template/static")), "1800")))
	http.ListenAndServe(":80", nil)
}

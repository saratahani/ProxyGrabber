package main

import (
	"bytes"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

func cleaner(url string) (string, error) {

	re, err := http.Get(url)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return "", err
	}

	r := strings.NewReader(string(b))

	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	body := cascadia.MustCompile("textarea").MatchFirst(doc)
	html.Render(buf, body)

	s := buf.String()

	m := strings.Replace(s, `<textarea onclick="this.focus();this.select()" style="font-size: 11pt; font-weight: bold; width: 500px; height: 300px; background-color: #000000; color: #0065dd;" wrap="hard">`, "", 1)

	m1 := strings.Replace(m, `</textarea>`, "", 1)

	return m1, err

}

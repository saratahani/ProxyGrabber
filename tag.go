package main

import (
	"golang.org/x/net/html"
	"net/http"
	"regexp"
)

func getTag() []string {
	var links []string

	r := regexp.MustCompile(`^http://www.socksproxylist24.top/\d.*html$`)

	re, _ := http.Get(`http://www.socksproxylist24.top/`)

	z := html.NewTokenizer(re.Body)

	for {

		tt := z.Next()

		switch {

		case tt == html.ErrorToken:
			return links
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"

			if isAnchor {

				for _, a := range t.Attr {

					if a.Key == "href" {

						if r.MatchString(a.Val) {
							links = append(links, a.Val)
						}
						break
					}
				}
			}
		}
	}
}

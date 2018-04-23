package main

import (
	"golang.org/x/net/html"
	"net/http"
	"regexp"
)

//Getting tags from site
func getTag() []string {
	var links []string

	r := regexp.MustCompile(`^http://www.socksproxylist24.top/\d.*html$`)

	re, _ := http.Get(`http://www.socksproxylist24.top/`)

	z := html.NewTokenizer(re.Body)

	for {

		switch token := z.Next(); token {

		//Return slice with links if tags end
		case html.ErrorToken:
			return links

		case html.StartTagToken:

			t := z.Token()

			//Check if is tag is anchor
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

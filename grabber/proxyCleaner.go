package grabber

import (
	"bytes"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

//cleaner removes textarea tags from proxies
func cleaner(url string) string {

	response, _ := http.Get(url)

	b, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	r := strings.NewReader(string(b))

	doc, _ := html.Parse(r)

	buf := new(bytes.Buffer)

	//Search for textarea tag
	body := cascadia.MustCompile("textarea").MatchFirst(doc)
	html.Render(buf, body)

	proxies := buf.String()

	cleanProxiesTop := strings.Replace(proxies, `<textarea onclick="this.focus();this.select()" style="font-size: 11pt; font-weight: bold; width: 500px; height: 300px; background-color: #000000; color: #0065dd;" wrap="hard">`, "", 1)

	cleanProxiesDown := strings.Replace(cleanProxiesTop, `</textarea>`, "", 1)

	return cleanProxiesDown

}

// martini.go for learning purpose
package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/codegangsta/martini"
	"io"
	"net/http"
	"sort"
	"strings"
)

var token = "haoyunmommy_1234567"

func main() {
	// Start the martini instance.
	m := martini.Classic()

	m.Get("/weixin", func(r *http.Request) string {
		query_url := r.URL.Query()
		signature, timestamp, nonce, echostr := query_url.Get("signature"), query_url.Get("timestamp"), query_url.Get("nonce"), query_url.Get("echostr")

		arr := []string{token, timestamp, nonce}
		sort.Strings(arr)
		// 		fmt.Println(arr)

		sha1String := weixin_sha1(strings.Join(arr, ""))
		if sha1String == signature {
			return echostr
		}
		return "not the right signature"
	})
	m.Run()
}

func weixin_sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

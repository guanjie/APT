// martini.go for learning purpose

package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/codegangsta/martini"
	"io"
)

func main() {
	m := martini.Classic()

	m.Get("/weixin", func(params martini.Params) string {
		fmt.Println(params["signature"])
		return params["signature"]
	})
	m.Run()
}

func weixin_sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

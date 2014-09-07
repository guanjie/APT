// main.go file for testing purpose.
package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()
	// render html templates from templates directory
	m.Use(martini.Static("assets/"))
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		//r.HTML(200, "index", "eric")
		r.HTML(200, "home", "eric")
	})

	// Adding the wxpay-jsapi-demo.html testing part
	m.Get("/weixin/pay", func(r render.Render) {
		//r.HTML(200, "index", "eric")
		r.HTML(200, "wxpay-jsapi-demo", "eric")
	})

	// Adding the wxpay-jsapi-demo.html testing part
	m.Get("/staging", func(r render.Render) {
		//r.HTML(200, "index", "eric")
		r.HTML(200, "staging_index", "eric")
	})

	m.Get("/api/getTopicList", func(r render.Render, req *http.Request) {
		fmt.Println("getTopicList")

		query_url := req.URL.Query()
		n := query_url.Get("n")
		fmt.Println(n)
		//r.JSON(200,"{'n':'32'}")
		newmap := map[string]interface{}{"responseText": "success"}
		r.JSON(200, newmap)
		//r.HTML(200, "index", "dddd")
		//r.Redirect("/")
	})

	m.Run()
}

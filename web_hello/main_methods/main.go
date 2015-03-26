// martini.go for learning purpose
package main

import (
	"github.com/codegangsta/martini"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	m := martini.Classic()
	m.Get("/eric", func(r *http.Request) string {
		return "Hello world"
	})

	// 算法开始
	m.Get("/eric/number/:num", func(params martini.Params) string {
		num, _ := strconv.Atoi(params["num"])
		num = powerto3(num)
		return "Hello the number " + params["num"] + " to the power 3 is: " + strconv.Itoa(num)
	})

	// 拿到一个城市的温度
	m.Get("/eric/weather/:city", func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		resp, _ := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + params["city"])
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	})

	// 英文字典
	m.Get("/eric/dictionary/:word", func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		resp, _ := http.Get("https://api.pearson.com/v2/dictionaries/ldec/entries?headword=" + params["word"])
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	})

	m.Run()
}

func powerto3(num int) int {
	return num * num * num
}

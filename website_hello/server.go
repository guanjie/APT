package main

import "Classic"
import "martini"

func main() {
    m := martini.Classic()
    m.Get("/", func() string {
        return "Check This out!"
    })
    m.Run()
}

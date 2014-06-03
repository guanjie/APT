// urlparser.go file, for learning purpose

package main

import (
    "fmt"
    "net/url"
)

var p = fmt.Println

func main() {
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    p(u)

    p(u.Scheme)
    p(u.User)
    p(u.User.Username())
    p(u.User.Password())

    p(u.Host)
    p(u.Path)
    p(u.Fragment)
    p(u.RawQuery)

}

package main

import "github.com/go-martini/martini"
import "github.com/Martini/WebHander/DB"

func main() {
	m := martini.Classic()
	m.Get("/", DB.ReT)
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})
	m.Run()
}

package main

import (
	"gopkg.in/macaron.v1"
)
func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/hello", func() (int, string) {
		return 400, "Hello world!"
	})
	m.Run()
}

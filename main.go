package main

import (
	"github.com/nesquikmike/small-web-ldn-bus-times/controllers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/countdown", c.Countdown)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

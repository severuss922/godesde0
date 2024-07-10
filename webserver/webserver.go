package webserver

import (
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/index", home)
	http.ListenAndServe(":8080", nil)
}

func home(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, "./webserver/index.html")
}

package main

import (
	"net/http"

	"../../library"
	broker "./src"
)

func main() {
	http.HandleFunc("/", broker.WSHandler)
	library.StartServer(":8000")
}

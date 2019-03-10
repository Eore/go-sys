package main

import (
	"log"
	"net/http"

	"../../library"
	broker "./src"
)

const port = ":8000"

func main() {
	http.HandleFunc("/", broker.WSHandler)
	log.Println("Server start in port" + port)
	library.StartServer(port)
}

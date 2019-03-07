package broker

import "github.com/gorilla/websocket"

type Message struct {
	Domain    string      `json:"domain"`
	Action    string      `json:"action"`
	Parameter interface{} `json:"parameter"`
}

type Pool struct {
	Clients []Client
}

type Client struct {
	Domain     string
	Connection *websocket.Conn
}

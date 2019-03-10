package broker

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Domain    string      `json:"domain,omitempty"`
	Type      string      `json:"type"`
	Action    string      `json:"action,omitempty"`
	Status    Status      `json:"status,omitempty"`
	Parameter interface{} `json:"parameter,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type Status struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Pool struct {
	Clients []Client
}

type Client struct {
	Domain     string
	IP         string
	Connection *websocket.Conn
}

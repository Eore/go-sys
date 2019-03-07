package broker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var m Message
var p Pool

func WSHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("%s connect", conn.RemoteAddr())

	go func(conn *websocket.Conn) {
		for {
			err := conn.ReadJSON(&m)
			if err != nil {
				log.Printf("%s disconnect", conn.RemoteAddr())
				conn.Close()
				return
			}
			switch m.Action {
			case "join":
				p.AddClient(Client{
					Domain:     m.Domain,
					Connection: conn,
				})
				log.Printf("%s joined\n", m.Domain)
			default:
				log.Printf("Sending to: %+v\n", m)
				p.SendMessage(m)
			}
		}
	}(conn)
}

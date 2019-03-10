package broker

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var message Message
var pool Pool

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
			err := conn.ReadJSON(&message)
			IP := conn.RemoteAddr().String()
			if err != nil {
				log.Printf("%s disconnect", IP)
				log.Printf("Remove %+v from list", IP)
				pool.RemoveClient(IP)
				conn.Close()
				return
			}
			switch message.Action {
			case "list":
				conn.WriteJSON(pool.ListClient())
				fmt.Println(pool.ListClient())
			case "join":
				client := Client{
					Domain:     message.Domain,
					IP:         IP,
					Connection: conn,
				}
				log.Printf("Adding %+v to list", client)
				_, err := pool.AddClient(client)
				if err != nil {
					log.Printf("%s already joined\n", client.Domain)
				} else {
					pool.SendMessage(Message{
						Type:   "response",
						Status: SUCCESS,
					})
					log.Printf("%s joined\n", message.Domain)
				}
			default:
				log.Printf("Sending to: %+v\n", message)
				pool.SendMessage(message)
			}
		}
	}(conn)
}

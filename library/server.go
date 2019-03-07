package library

import (
	"log"
	"net/http"
)

// func WsHandler(w http.ResponseWriter, r *http.Request) {
// 	upgrader := websocket.Upgrader{}

// 	wsCon, _ := upgrader.Upgrade(w, r, nil)

// 	for {
// 		messageType, p, err := wsCon.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		if err := wsCon.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

//StartServer (starting server)
func StartServer(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}

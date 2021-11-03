package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func send(conn *websocket.Conn)  {
	for  {
		time.Sleep(1 * time.Second)
		t := time.Now().Format("15:04:05")
		err := conn.WriteMessage(1, []byte(t))
		if err != nil {
		}
	}

}

func timer(w http.ResponseWriter, r *http.Request)  {
	upgrade.CheckOrigin = func(r *http.Request) bool {return true}
	ws, err := upgrade.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Клиент подсоединился")
	send(ws)


}

func main() {

	//fmt.Println(time.Now().Format("15:04:05"))

	http.HandleFunc("/time", timer )
	err := http.ListenAndServe(":7001", nil)
	if err != nil {
		 log.Println(err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func send(conn *websocket.Conn) {
	for {
		time.Sleep(1 * time.Second)
		t := time.Now().Format("15:04:05")
		err := conn.WriteMessage(1, []byte(t))
		if err != nil {
		}
	}

}

func timer(w http.ResponseWriter, r *http.Request) {
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Клиент подсоединился")
	send(ws)
}

func temp(w http.ResponseWriter, r *http.Request) {
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	tmp, _ := template.ParseFiles("index.html")
	err := tmp.Execute(w, "")
	if err != nil {
		log.Println("Ошибка")
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "twitter.ico")
}

func main() {
	http.HandleFunc("/", temp)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/time", timer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thinkerou/favicon"
	"html/template"
	"log"
	"net/http"
	"time"
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

func timer(c *gin.Context) {
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Клиент подсоединился")
	send(ws)
}

func temp(c *gin.Context) {
	tmp, _ := template.ParseFiles("index.html")
	err := tmp.Execute(c.Writer, "")
	if err != nil {
		log.Println("Ошибка")
	}
}

func main() {
	server := gin.Default()
	server.Static("/css", "css")
	server.Use(favicon.New("./favicon.ico"))
	server.GET("/", temp)
	server.GET("/time", timer)

	err := server.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}

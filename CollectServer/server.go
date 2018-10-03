package main

import (
	"./conf"
	"./router"
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"net/http"
)

func main() {
	go func() {
		a()
	}()

	if err := conf.Init(""); err == nil {
		fmt.Println("config success")
	}
	router.RunSubDomains()

}

func a() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		fmt.Println("on connection")
		//so.Join("chat")
		//so.On("chat message", func(msg string) {
		//	log.Println("emit:", so.Emit("chat message", msg))
		//	so.BroadcastTo("chat", "chat message", msg)
		//})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
		so.On("reqSignIn", func(data map[string]string) {
			log.Println(data)
			http.Get("http://localhost:3000/api/v1/sign-in")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		server.ServeHTTP(w, r)
	})
	//http.Handle("/socket.io/", server)
	//http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:2100...")
	log.Fatal(http.ListenAndServe(":2100", nil))
}

func b() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	//handle connected
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		//join them to room
		c.Join("chat")
	})

	type Message struct {
		Name string `json:"name"`
		Message string `json:"message"`
	}

	//handle custom event
	server.On("send", func(c *gosocketio.Channel, msg Message) string {
		//send event to all in room
		c.BroadcastTo("chat", "message", msg)
		return "OK"
	})

	//setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	log.Panic(http.ListenAndServe(":3100", serveMux))
}

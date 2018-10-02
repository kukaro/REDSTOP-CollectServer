package main

import (
	"./conf"
	"./router"
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	go func() {
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
				log.Println(data["id"])
			})
		})
		server.On("error", func(so socketio.Socket, err error) {
			log.Println("error:", err)
		})
		http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			server.ServeHTTP(w, r)
		})
		//http.Handle("/socket.io/", server)
		//http.Handle("/", http.FileServer(http.Dir("./asset")))
		log.Println("Serving at localhost:3100...")
		log.Fatal(http.ListenAndServe(":3100", nil))
	}()

	go func(){
		if err := conf.Init(""); err == nil {
			fmt.Println("config success")
		}
		router.RunSubDomains()
	}()
}

package main

import (
	"log"
	"fmt"
	s "github.com/PyMarcus/message_queue/server"
	st "github.com/PyMarcus/message_queue/storage"
	
	"github.com/gorilla/websocket"
)

func main(){
	fmt.Println("->STARTING MESSAGE QUEUE<-")
	
	cfg := &s.Config{
		ListenAddr: ":7777",
		WebSocketAddr: ":6666",
		StorageProducer: st.NewMemoryStore,
	}
	server, error := s.NewServer(cfg)
	if error != nil{
		panic(error)
	}
	consumerWSConn()

	server.RunAndListen()
	select{}
}

func consumerWSConn(){
  conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:6666", nil)
  if err != nil{
    log.Fatal(err)
  }
  
  log.Println(conn)
}
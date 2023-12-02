package main

import (
	"encoding/json"
	"fmt"
	"log"

	s "github.com/PyMarcus/message_queue/common"
	st "github.com/PyMarcus/message_queue/storage"

	"github.com/gorilla/websocket"
)

func main(){
	fmt.Println("==========STARTING MESSAGE QUEUE==========")
	
	cfg := &s.Config{
		ListenAddr: ":7777",
		WebSocketAddr: ":6666",
		StorageProducer: st.NewMemoryStore,
	}
	server, error := s.NewServer(cfg)
	if error != nil{
		panic(error)
	}
	// this is for tests -> firts, comment this to execute server ,and, then, uncomment this to make requests over websockets.
	// npm install -g wscat
	// wscat -c ws://localhost:6666
	//consumerWSConn()
	server.RunAndListen()
	select{}
}

func consumerWSConn(){
  conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:6666", nil)
  if err != nil{
    log.Fatal(err)
  }
  msg := &s.WSMessage{Action: "subscribe", Topic: "vanilla"}
  log.Println("Sending message ", msg)
  data, err := json.Marshal(msg)
  if err != nil{
     log.Println("Fail to marshal message")
     log.Fatal(err)
  }
  conn.WriteMessage(websocket.TextMessage, data)
}
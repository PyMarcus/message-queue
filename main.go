package main

import (
	"fmt"
	s "github.com/PyMarcus/message_queue/server"
	st "github.com/PyMarcus/message_queue/storage"
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
	server.RunAndListen()
	select{}
}
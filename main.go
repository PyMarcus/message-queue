package main 

import (
	"fmt"
	s "github.com/PyMarcus/message_queue/server"
	st "github.com/PyMarcus/message_queue/storage"
)

func main(){
	fmt.Println("starting queue")
	
	cfg := &s.Config{
		ListenAddr: ":7777",
		StorageProducer: st.NewMemoryStore,
	}
	_, error := s.NewServer(cfg)
	if error != nil{
		panic(error)
	}

	
}
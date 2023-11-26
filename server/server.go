package server

import (
	"log"
	//"net/http"

	m "github.com/PyMarcus/message_queue/message"
	st "github.com/PyMarcus/message_queue/storage"
	tr "github.com/PyMarcus/message_queue/transport"
)

type Config struct{
	ListenAddr string
	StorageProducer st.StorageProducer
}

type Server struct{
	*Config
	topics map[string]st.Storage
	consumers []tr.Consumer
	producers []tr.Producer
	producersCh chan m.Message
	quitch    chan struct{}
}

// base functions
func NewServer(cfg *Config) (*Server, error){
    pm := make(chan m.Message)

	return &Server{
		Config: cfg,
		topics: make(map[string]st.Storage),
		quitch: make(chan struct{}),
		producers: []tr.Producer{tr.NewHTTPProducer(cfg.ListenAddr, pm)},
		producersCh: pm,
	}, nil
}

func (s *Server) RunAndListen(){	
    for _, consumer := range s.consumers{
       if err := consumer.Start(); err != nil{
           log.Println(err)
           continue
       }
    }
    
    for _, producer := range s.producers{
	 go func(p tr.Producer){
		if err := p.Start(); err != nil{
			log.Println(err)
			return
		}
	  }(producer)
	 }		 
	 s.loop()
}

func (s Server) createTopicIfNotExists(name string) st.Storage{
	_, exists := s.topics[name]
	
	if !exists{
		s.topics[name] = &st.MemoryStore{}
		log.Println("::TOPIC CREATED -> ", name)
	}	
	return s.topics[name]
}

func (s *Server) publish(message m.Message) (int, error){
    store := s.createTopicIfNotExists(message.Topic)
	return store.Push(message.Data)
}

func (s *Server) loop(){
   for{
      select{
	  case <-s.quitch:
	      return
	  case msg := <- s.producersCh:
	      if offset, err := s.publish(msg); err != nil{
	        log.Println("XX ERROR TO PUBLISH ", err)
	      }else{
			log.Printf("::PRODUCER RECEIVED +%d DATA ON -> %s\n\n\n", offset + 1, msg.Topic)
	      }
      }   
   }
}
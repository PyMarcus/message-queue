package server

import (
	"log"
	//"net/http"

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
	quitch    chan struct{}
}

// base functions
func NewServer(cfg *Config) (*Server, error){
	return &Server{
		Config: cfg,
		topics: make(map[string]st.Storage),
		quitch: make(chan struct{}),
		producers: []tr.Producer{tr.NewHTTPProducer(cfg.ListenAddr)},
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
		if err := producer.Start(); err != nil{
			log.Println(err)
			continue
		}
	 }
	<- s.quitch
}

func (s Server) createTopic(name string) bool{
	_, exists := s.topics[name]
	
	if !exists{
		s.topics[name] = &st.MemoryStore{}
		return true
	}
	
	return false
}


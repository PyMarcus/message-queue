package server

import (
	"net/http"

	st "github.com/PyMarcus/message_queue/storage"
)

type Config struct{
	ListenAddr string
	StorageProducer st.StorageProducer
}

type Server struct{
	*Config
	topics map[string]st.Storage
}

func NewServer(cfg *Config) (*Server, error){
	return &Server{
		Config: cfg,
		topics: make(map[string]st.Storage),
	}, nil
}

func (s Server) RunAndListen(){
	http.ListenAndServe(s.Config.ListenAddr, nil)
}

func (s Server) createTopic(name string){
	_, exists := s.topics[name]
	
	if !exists{
		s.topics[name] = &st.MemoryStore{}
	}
}


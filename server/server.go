package server

import (
	"log"
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

//handlers
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Println(r.URL.Path)
}

// base functions
func NewServer(cfg *Config) (*Server, error){
	return &Server{
		Config: cfg,
		topics: make(map[string]st.Storage),
	}, nil
}

func (s *Server) RunAndListen(){	
	log.Fatal(http.ListenAndServe(s.Config.ListenAddr, s))
}

func (s Server) createTopic(name string) bool{
	_, exists := s.topics[name]
	
	if !exists{
		s.topics[name] = &st.MemoryStore{}
		return true
	}
	
	return false
}


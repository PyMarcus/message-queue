package common

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	m "github.com/PyMarcus/message_queue/message"
	st "github.com/PyMarcus/message_queue/storage"
	tr "github.com/PyMarcus/message_queue/transport"
	"github.com/gorilla/websocket"
)

var ppeers = make(map[string][]*WSPeer)

type Config struct{
	ListenAddr      string
	WebSocketAddr   string 
	StorageProducer st.StorageProducer
}

type Server struct{
	*Config
	mu     sync.RWMutex
	peers map[Peer]bool 
	topics    map[string]st.Storage
	consumers []Consumer
	producers []tr.Producer
	producersCh chan m.Message
	quitch      chan struct{}
}

// base functions
func NewServer(cfg *Config) (*Server, error){
    pm := make(chan m.Message)
    consumer, _ := NewWSConsumer(cfg.WebSocketAddr, &Server{})
	s := &Server{
		Config: cfg,
		peers: make(map[Peer]bool),
		topics: make(map[string]st.Storage),
		quitch: make(chan struct{}),
		producers: []tr.Producer{tr.NewHTTPProducer(cfg.ListenAddr, pm)},
		producersCh: pm,
		}
	s.consumers = append(s.consumers,  consumer,)
	return s, nil
}

func (s *Server) RunAndListen(){	
    for _, consumer := range s.consumers{
       go func(consumer Consumer){
		if err := consumer.Start(); err != nil{
			log.Println(err)
		}       
       }(consumer)
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

func (s *Server) createTopicIfNotExists(name string) st.Storage{
	_, exists := s.topics[name]
	
	if !exists{
		s.topics[name] = &st.MemoryStore{}
		log.Println("::TOPIC CREATED -> ", name)
		if _, ok := ppeers[name]; !ok {

			ppeers[name] = []*WSPeer{}
		}else{
			ppeers[name] = append(ppeers[name], ppeers[name]...)
		}
	
	
		log.Println("creating peers's store -> ", ppeers)
	}	
	return s.topics[name]
}

func (s *Server) publish(message m.Message) (int, error) {
    store := s.createTopicIfNotExists(message.Topic)
    return store.Push([]byte(message.Data))
}

func (s *Server) loop(){
   for{
      select{
	  case <-s.quitch:
	      return
	  case msg := <- s.producersCh:
		  s.mu.Lock()
		  offset, err := s.publish(msg)
		  s.mu.Unlock()
	      if err != nil{
	        log.Println("ERROR TO PUBLISH ", err)
	      }else{
			log.Printf("PRODUCER POST +%d DATA ON -> %s\n\n\n", offset + 1, msg.Topic)
			peers := ppeers[strings.TrimSpace(msg.Topic)]
			if len(ppeers) > 0{
				for i, p := range peers{
				   if i == offset + 1{
					  store := s.createTopicIfNotExists(msg.Topic)
				      store.ClearMemory()
				      break
				   }
				   if p != nil{
					 p.Send([]byte(msg.Data))
				   }
				}
		    }
	      }
      }         
   }
}

func (s *Server) AddPeer(conn Peer){
    s.mu.Lock()
	defer s.mu.Unlock()
	
    if s.peers == nil {
        s.peers = make(map[Peer]bool)
    }
	s.peers[conn] = true
	log.Println("Added new peer ", conn)
}

func (s *Server) AddPeerToTopic(topic string, peer *WSPeer){
	s.mu.Lock()

	defer s.mu.Unlock()
	
	if peer != nil {
	    if peers, ok := ppeers[topic]; ok {
			ppeers[topic] = append(peers, peer)
		} else {
		    ppeers = make(map[string][]*WSPeer)
			log.Println("Adding peer to topic -> ", topic, peer)
			ppeers[topic] = []*WSPeer{peer}
			log.Println(" PEERS ", len(ppeers[topic]))
		}
		
    }
}

func (s *Server) notifyConsumerByTopic(topic, data string){

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:6666", nil)
	if err != nil{
	    log.Fatal(err)
	}
	
	msg := &m.Message{Data: data, Topic: topic}
	log.Println("Sending message: ", msg.Data, " to consumers into topic: ", topic)

	content, err := json.Marshal(msg)
	if err != nil{
	     log.Println("Fail to marshal message")
	     log.Fatal(err)
	}
	conn.WriteMessage(websocket.TextMessage, content)
}

func (s *Server) RemovePeer(peer *WSPeer, topic string){
    log.Println("Peer was removed!")
    index := 0
	for i, v := range ppeers[topic]{
	  if v == peer{
	     index = i 
	  }
	}
	ppeers[topic] = append(ppeers[topic][:index], ppeers[topic][index+1:]...)
}
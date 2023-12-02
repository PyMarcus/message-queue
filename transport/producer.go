package transport

import (
	"log"
	"net/http"
	"strings"
	"io/ioutil"

	m "github.com/PyMarcus/message_queue/message"
)

type Producer interface{
   Start() error 
}

type HTTPProducer struct{
    Addr string 
    producerCh chan <- m.Message  // write channel
}

func NewHTTPProducer(listenAddr string, producerCh chan m.Message) Producer{
    return &HTTPProducer{
        Addr: listenAddr, 
        producerCh: producerCh,
    }
}

func (h *HTTPProducer) Start() error{
	log.Println("Producer started on ", h.Addr)
    return http.ListenAndServe(h.Addr, h)
}

//handlers
func (h *HTTPProducer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	path := strings.TrimPrefix(r.URL.Path, "/") 
	parts := strings.Split(path, "/")
	
	if r.Method == "GET"{}
	
	if r.Method == "POST"{
	    if len(parts) != 2{
	        w.Write([]byte("Bad request!"))
	        return
		}
		
		topic := parts[1]
		
		log.Println("::TOPIC -> ", topic)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{
			w.Write([]byte("Bad request!"))
	        return
		}
		h.producerCh <- m.Message{
		   Topic: topic,
		   Data: string(body),
		}
		w.Write([]byte("Added!"))
	}
}

func (h *HTTPProducer) handlePublish(){
      
}

package transport

import (
	"log"
	"net/http"
	"strings"
)

type Consumer interface{
   Start() error 
}

type Producer interface{
   Start() error 
}

type HTTPProducer struct{
    Addr string 
}

func NewHTTPProducer(listenAddr string) Producer{
    return &HTTPProducer{
        Addr: listenAddr, 
    }
}

func (h *HTTPProducer) Start() error{
	log.Println("HTTP Transport started on ", h.Addr)
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
		log.Println(topic)
	}
}

func (h *HTTPProducer) handlePublish(){
      
}

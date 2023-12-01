package transport

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var updgrade = websocket.Upgrader{}

func req(){
	websocket.DefaultDialer.Dial("ws:/oo", nil)
	
}

type Consumer interface{
	Start() error 
}
 
type WSConsumer struct{
    
}

func (ws *WSConsumer) Start() error{
    return http.ListenAndServe(":4000", ws)
}

func (ws *WSConsumer) ServeHTTP(w http.ResponseWriter, r *http.Request){
    updgrade.Upgrade(w, r, nil)
}

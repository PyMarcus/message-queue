package server

import (
	"log"
	"net/http"

	"github.com/PyMarcus/message_queue/peer"
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
	ListenAddr string
	server *Server
	peers  chan *websocket.Conn
}

func NewWSConsumer(address string, serv *Server) (*WSConsumer, error){
    return &WSConsumer{ListenAddr: address, server: serv}, nil
}

func (ws *WSConsumer) Start() error{
	log.Println("Consumer Listening on ", ws.ListenAddr)
    return http.ListenAndServe(ws.ListenAddr, ws)
}

func (ws *WSConsumer) ServeHTTP(w http.ResponseWriter, r *http.Request){
    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true // Allow all origins for testing purposes
        },
    }
    
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil{
       log.Println("fail to connect into websocket! ", err)
       return
    }
    p := peer.NewPeer(conn)
    ws.server.AddPeer(p)
}

type WSMessage struct{
    Topic  string  `json:"topic"`
    Action string  `json:"action"`
}

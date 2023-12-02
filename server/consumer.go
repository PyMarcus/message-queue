package server

import (
	"log"
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
	ListenAddr string
	server *Server
	peers  chan *websocket.Conn
}

func NewWSConsumer(address string) (*WSConsumer, error){
    return &WSConsumer{ListenAddr: address}, nil
}

func (ws *WSConsumer) Start() error{
	log.Println("Consumer Listening on ", ws.ListenAddr)
    return http.ListenAndServe(ws.ListenAddr, ws)
}

func (ws *WSConsumer) ServeHTTP(w http.ResponseWriter, r *http.Request){
    conn, err := updgrade.Upgrade(w, r, nil)
    if err != nil{
       log.Println("fail to connect into websocket! ", err)
       return
    }
        
    ws.server.AddPeer(conn)
    log.Println(conn)
}

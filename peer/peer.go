package peer

import (
    "log"
	"github.com/PyMarcus/message_queue/message"
	"github.com/gorilla/websocket"
)

type Peer interface{
   Send([]byte) error
}

type WSPeer struct{
    conn *websocket.Conn
}

//:::MAYBE, this could to produce memory leaks!
func NewPeer(conn *websocket.Conn) *WSPeer{
    p := &WSPeer{
        conn: conn,
    }
    
    // read loop
    go func(){
       var msg message.Message
       for{
          if err := p.conn.ReadJSON(&msg); err != nil{
             log.Println("fail to read message in peer loop! ", err)
             return
          }
          
          if err := p.handleMessage(msg); err != nil{
             log.Println("Fail to parse message", err)
             return
          }
       }
    }()
    return p
}

func (p *WSPeer) Send(b []byte) (err error){
    err = p.conn.WriteMessage(websocket.BinaryMessage, b)
    return
}

func (p *WSPeer) handleMessage(msg message.Message) error{
   log.Println("handling message", msg)  
   return nil   
}
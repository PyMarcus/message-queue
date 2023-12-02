package peer

import (
	"github.com/gorilla/websocket"
)

type Peer interface{
   Send([]byte) error
}

type WSPeer struct{
    conn *websocket.Conn
}

func NewPeer(conn *websocket.Conn) *WSPeer{
    return &WSPeer{
        conn: conn,
    }
}

func (p *WSPeer) Send(b []byte) (err error){
    err = p.conn.WriteMessage(websocket.BinaryMessage, b)
    return
}
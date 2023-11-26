package transport

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestPublisher(t *testing.T){
    topics := []string{"topicA", "topicB", "topicC"}
    for i, topic := range topics{
        payload := []byte("my message to test " + strconv.Itoa(i))

       response, err := http.Post("http://localhost:7777/publish/" + topic, "application/octet-stream", bytes.NewReader(payload))
       if err != nil{
          t.Error(err)
       }
       body, _ := io.ReadAll(response.Body)
       log.Println(body)
   }
}
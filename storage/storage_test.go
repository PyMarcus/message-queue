package storage

import (
	"fmt"
	"testing"
)


// test push and fetch
func TestStorage(t *testing.T){
	storage := NewMemoryStore()
	for i := 1; i < 11; i++{
		offset, err := storage.Push([]byte(fmt.Sprintf("Test %d", i)))
		if err != nil{
			t.Errorf("error %s", err)
		}
		
		data, err := storage.Fetch(offset)
		if err != nil{
			t.Errorf("error2 %s", err)
		}
		
		fmt.Println(string(data))
	}
	
}
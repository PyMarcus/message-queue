# message-queue


### RUN:

    make 


### Use:

     * Publisher
    
     curl -X POST http://localhost:7777/publish/topicname --data-binary 'some message'
     
     
     * Consumer (websocket)
     
     ws://localhost:6666  
     
     
     json body:
       
       
        {
	        "topic":"vanilla",
	        "data":"foda-se"
        }
     
# message-queue


![image](https://github.com/PyMarcus/message-queue/assets/88283829/941b403e-194f-4dea-ad13-f3ccab86603e)



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


### Publish:

![image](https://github.com/PyMarcus/message-queue/assets/88283829/5a705e66-4f1f-4c9f-9f68-59648c100a8f)


### Consumer:

![image](https://github.com/PyMarcus/message-queue/assets/88283829/da43b5f2-ac2f-4407-b51c-1d2fb82da488)


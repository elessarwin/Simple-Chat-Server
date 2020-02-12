##How it works?
Receiving messages will send to all node.
## Client Pattern
send message to server-address:8080 whit this struct:

    message ChatMessage{
        string UserId = 1 ;
        string Message = 2 ;
    }
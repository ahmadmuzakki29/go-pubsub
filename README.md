go-pubsub

Golang implementation of Google Cloud Pub/Sub

###Installation

`go get -u github.com/ahmadmuzakki29/go-pubsub`


###Usage

```
import "github.com/ahmadmuzakki29/go-pubsub"
func main(){
    myProjectId := "YOUR PROJECT ID GOES HERE"
    pubsub.Init({ProjectId:myProjectId})
    
    msg := `{"name":"John"}`
    err := Publish("hello", []byte(msg))
    if err != nil {
        t.Error(err)
    }
    
    AddHandler("hello","helloChannel",handler)
}

func handler(msg *Message)error{
    // do something with your message
    // --e.g msg.ToJson(&dest)
    msg.Done(true)
    return nil
}
```
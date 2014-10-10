package main

import "fmt"

func main() {

    // Create a new channel
    // Channels are typed by the values they send.
    messages := make(chan string)

    // _Send_ a value into a channel using the `channel <-`
    // syntax. Here we fire off a goroutine that sends "ping"
    // to the messages channel.
    go func() { messages <- "ping" }()
   
    // Here we'll receive the `"ping"` message
    // we sent above and print it out.
    msg := <-messages
    fmt.Println(msg)
    // TODO describe select BETTER
//  select {
//  case msg1 := <- c1:
//    fmt.Println("Message 1", msg1)
//  case msg2 := <- c2:
//    fmt.Println("Message 2", msg2)
//  case <- time.After(time.Second):
//    fmt.Println("timeout")
    }
}

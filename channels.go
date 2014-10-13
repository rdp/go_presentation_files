package main

import "fmt"

func main() {

    // Create a new channel to send messages of type "string"
    messages := make(chan string)

    // _Send_ a value into a channel using the `channel <-`
    go func() { messages <- "ping" }()
   
    // Here we'll wait for the "ping" message
    msg := <-messages
    fmt.Println("received via channel:" + msg)
}

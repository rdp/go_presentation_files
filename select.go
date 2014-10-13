package main

import "fmt"
import "time"

func main() {

    messages := make(chan string)
    messages2 := make(chan string)

    go func() { messages <- "ping" }()
    go func() { messages2 <- "ping" }()
   
    select {
    case msg1 := <- messages:
      fmt.Println("Message 1", msg1)
    case msg2 := <- messages2:
      fmt.Println("Message 2", msg2)
    case <- time.After(time.Second):
      fmt.Println("timeout")
    }
    fmt.Println("done")
}


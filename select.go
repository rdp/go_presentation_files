// uncompilable
    
  select {
  case msg1 := <- c1:
    fmt.Println("Message 1", msg1)
  case msg2 := <- c2:
    fmt.Println("Message 2", msg2)
  case <- time.After(time.Second):
    fmt.Println("timeout")
  }

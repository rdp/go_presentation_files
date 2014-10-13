package main

import "fmt" 
import "time"        

func ourMethod(input string) {
   fmt.Println(input)
}

func main() {
  ourMethod("non goroutine")
  go ourMethod("goroutine")
  go ourMethod("goroutine2")
  ourMethod("non goroutine again")
  
  time.Sleep(100 * time.Millisecond) // make sure goroutines finish
}
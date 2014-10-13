package main

import "fmt" 
import "time"        

func f(input string) {
   fmt.Println(input)
}

func main() {
  f("non goroutine")
  go f("goroutine")
  go f("goroutine2")
  
  time.Sleep(100 * time.Millisecond) // make sure goroutines finish
}
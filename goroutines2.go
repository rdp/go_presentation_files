package main

import "fmt" 
import "time"        
//import "runtime" 

func f(input string) {
    for i := 0; i < 30000; i++ {
        fmt.Println(input, ":", i)
    }
}

func main() {
  //runtime.GOMAXPROCS(2)
  go f("goroutine")
  go f("goroutine2")
  f("non goroutine")
  
  time.Sleep(100 * time.Millisecond) // let goroutines finish
}
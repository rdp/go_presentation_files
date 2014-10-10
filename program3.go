package main

import "fmt" 
import "time"        
import "runtime" 

func f(input string) {
    for i := 0; i < 300; i++ {
        fmt.Println(input, ":", i)
    }
}

func main() {
  runtime.GOMAXPROCS(2)
  f("non goroutine")
  go f("goroutine")
  go f("goroutine2")
  
  time.Sleep(100 * time.Millisecond) // let goroutine finish
}
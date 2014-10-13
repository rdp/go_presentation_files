package main

import "fmt" 
import "time"        

func f(input string) {
    for i := 0; i < 30000; i++ {
        fmt.Println(input, ":", i)
    }
}

func main() {
  go f("goroutine")
  f("non goroutine")
  
  time.Sleep(100 * time.Millisecond) // let goroutines finish
}
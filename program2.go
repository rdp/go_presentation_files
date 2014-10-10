package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
    _, err := fmt.Println("what if this fails?")
    fmt.Printf("did it fail %v?", bytesWritten, err)
}
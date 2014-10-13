package main

import "fmt"

func main() {
    bytesWritten, err := fmt.Println("what if this fails?")
    fmt.Printf("did it fail %v?", bytesWritten, err)
    // bytesWritten, _ := fmt.Println("what if this fails?")
}
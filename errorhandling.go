package main

import "fmt"

func main() {
    bytesWritten, err := fmt.Println("what if this fails?")
    fmt.Printf("did it fail bytesWritten= %v err=%v?", bytesWritten, err)
    bytesWritten, _ = fmt.Println("")
    fmt.Printf("bytesWritten = %d", bytesWritten)
}